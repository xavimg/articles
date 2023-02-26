package models

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/xavimg/articles/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/juju/errors"
)

var Repo *Database

func ConnectRepo(ctx context.Context) error {
	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.Settings.Mongo.URL))
	if err != nil {
		return errors.Trace(err)
	}

	Repo = &Database{
		mongo: c.Database(config.Settings.Mongo.Database),
	}

	logrus.Info("Connected to mongo!")

	return nil
}

type Database struct {
	mongo *mongo.Database
}

func (repo *Database) Articles() *mongo.Collection {
	return repo.mongo.Collection("articles")
}

func (repo *Database) List(ctx context.Context, team string) ([]*Article, error) {
	cursor, err := Repo.Articles().Find(ctx, bson.M{"teamId": team})
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

	if len(articles) == 0 {
		return nil, errors.Trace(errors.NotFoundf("team %s", team))
	}

	return articles, nil
}

func (repo *Database) Get(ctx context.Context, team, id string) (*Article, error) {
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

func (repo *Database) BatchCreate(ctx context.Context, articles []*Article) error {
	externalIds := []string{}
	for _, article := range articles {
		externalIds = append(externalIds, article.NewsArticleID)
	}

	// All existent
	cursor, err := Repo.Articles().Find(ctx, bson.M{"newsArticleID": bson.M{"$in": externalIds}})
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return errors.Trace(err)
	}
	defer cursor.Close(ctx)

	existingArticles := []*Article{}
	for cursor.Next(ctx) {
		article := &Article{}
		if err := cursor.Decode(&article); err != nil {
			return errors.Trace(err)
		}
		existingArticles = append(existingArticles, article)
	}
	if err := cursor.Err(); err != nil {
		return errors.Trace(err)
	}

	updates := []interface{}{}
	creates := []interface{}{}
	for _, article := range articles {
		if a := findArticle(existingArticles, article); a != nil {
			if !a.LastUpdateDate.Equal(article.LastUpdateDate) {
				updates = append(updates, article)
			}
			continue
		}
		creates = append(creates, article)
	}

	if len(updates) > 0 {
		if _, err = Repo.Articles().UpdateOne(ctx, bson.M{"newsArticleID": bson.M{"$in": externalIds}}, updates); err != nil {
			return errors.Trace(err)
		}
	}
	if len(creates) > 0 {
		if _, err = Repo.Articles().InsertMany(ctx, creates); err != nil {
			return errors.Trace(err)
		}
	}

	return nil
}

func findArticle(articles []*Article, article *Article) *Article {
	for _, a := range articles {
		if a.NewsArticleID == article.NewsArticleID {
			return a
		}
	}

	return nil
}
