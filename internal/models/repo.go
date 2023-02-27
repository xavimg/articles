package models

import (
	"context"

	"github.com/juju/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/xavimg/articles/internal/config"
)

var Repo *Database

// ConnectRepo will connect our server to any DB we opened the connection. Example: PostgreSQL snippet code
func ConnectRepo(ctx context.Context) error {
	mongo, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.Settings.Mongo.URL))
	if err != nil {
		return errors.Trace(err)
	}

	// example:
	//postgres, err := postgres.Open(...)

	Repo = &Database{
		mongo: mongo.Database(config.Settings.Mongo.Database),

		// example:
		// postgres: postgres
	}

	logrus.Info("Connected to mongo!")

	return nil
}

type Database struct {
	mongo *mongo.Database

	// postgres *postgres.DB
}

func (repo *Database) Articles() *mongo.Collection {
	return repo.mongo.Collection("articles")
}

func (repo *Database) GetArticle(ctx context.Context, team, id string) (*Article, error) {
	articleID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.Trace(err)
	}

	article := &Article{}
	filters := bson.M{"teamId": team, "_id": articleID}
	if err := Repo.Articles().FindOne(ctx, filters).Decode(&article); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.Trace(errors.NotFoundf(err.Error()))
		}
		return nil, errors.Trace(err)
	}

	return article, nil
}

func (repo *Database) ListArticles(ctx context.Context, team string) ([]*Article, error) {
	order := options.Find().SetSort(bson.M{"lastUpdateDate": -1})
	articles, err := repo.findArticles(ctx, bson.M{"teamId": team}, order)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if len(articles) == 0 {
		return nil, errors.Trace(errors.NotFoundf("team %s", team))
	}

	return articles, nil
}

func (repo *Database) ListArticlesByIds(ctx context.Context, team string, ids []string) ([]*Article, error) {
	order := options.Find().SetSort(bson.M{"lastUpdateDate": -1})
	articles, err := repo.findArticles(ctx, bson.M{"teamId": team, "newsArticleID": bson.M{"$in": ids}}, order)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return articles, nil
}

func (repo *Database) ReplaceOne(ctx context.Context, updates []*Article) error {
	for _, article := range updates {
		if _, err := Repo.Articles().ReplaceOne(ctx, bson.M{"newsArticleID": article.NewsArticleID}, article); err != nil {
			return errors.Trace(err)
		}
	}

	return nil
}

func (repo *Database) InsertMany(ctx context.Context, creates []interface{}) error {
	if _, err := Repo.Articles().InsertMany(ctx, creates); err != nil {
		return errors.Trace(err)
	}

	return nil
}

func (repo *Database) findArticles(ctx context.Context, filter interface{}, order *options.FindOptions) ([]*Article, error) {
	cursor, err := Repo.Articles().Find(ctx, filter, order)
	if err != nil {
		return nil, errors.Trace(err)
	}

	articles := []*Article{}
	for cursor.Next(ctx) {
		var article *Article
		if err := cursor.Decode(&article); err != nil {
			return nil, errors.Trace(err)
		}
		articles = append(articles, article)
	}
	if err := cursor.Err(); err != nil {
		return nil, errors.Trace(err)
	}

	return articles, nil
}
