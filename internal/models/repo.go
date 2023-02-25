package models

import (
	"context"
	"log"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoURL = "mongodb://mongo:27017"
)

var Repo *Database

func ConnectRepo(ctx context.Context) error {
	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURL))
	if err != nil {
		logrus.Errorf("Error connecting: %s\n", err)
		return err
	}

	Repo = &Database{
		mongo: c.Database("articles"),
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

func (repo *Database) GetAll(ctx context.Context) ([]*Article, error) {
	cursor, err := Repo.Articles().Find(ctx, bson.D{})
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	var articles []*Article
	for cursor.Next(context.Background()) {
		var article *Article
		if err := cursor.Decode(&article); err != nil {
			logrus.Error(err)
			return nil, err
		}
		articles = append(articles, article)
	}
	if err := cursor.Err(); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return articles, nil
}

func (repo *Database) GetByID(ctx context.Context, id string) (*Article, error) {
	var article Article

	articleID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	if err := Repo.Articles().FindOne(ctx, bson.M{"_id": articleID}).Decode(&article); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return &article, nil
}

func (repo *Database) InsertManyTask(ctx context.Context, articles []*Article) error {
	cursor, err := Repo.Articles().Find(ctx, bson.M{})
	if err != nil {
		logrus.Error(err)
		return err
	}
	defer cursor.Close(ctx)

	var existingArticles []Article
	for cursor.Next(ctx) {
		var article Article
		if err := cursor.Decode(&article); err != nil {
			logrus.Error(err)
			return err
		}
		existingArticles = append(existingArticles, article)
	}
	if err := cursor.Err(); err != nil {
		logrus.Error(err)
		return err
	}

	var newArticles []interface{}
	for _, article := range articles {
		if !containsArticle(existingArticles, article) {
			newArticles = append(newArticles, article)
		}
	}

	if len(newArticles) > 0 {
		_, err = Repo.Articles().InsertMany(ctx, newArticles)
		if err != nil {
			logrus.Error(err)
			return err
		}
	}

	return nil
}

func containsArticle(articles []Article, article *Article) bool {
	for _, a := range articles {
		if a.NewsArticleID == article.NewsArticleID {
			return true
		}
	}
	return false
}

func (repo *Database) InsertOne(ctx context.Context, article *Article) error {
	if _, err := Repo.Articles().InsertOne(ctx, article); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

// func fromModel(in *models.Article) *models.Article {
// 	return &Article{
// 		ArticleURL:        in.ArticleURL,
// 		NewsArticleID:     in.NewsArticleID,
// 		PublishDate:       in.PublishDate,
// 		Taxonomies:        in.Taxonomies,
// 		TeaserText:        in.TeaserText,
// 		ThumbnailImageURL: in.ThumbnailImageURL,
// 		Title:             in.Title,
// 		OptaMatchId:       in.OptaMatchId,
// 		LastUpdateDate:    in.LastUpdateDate,
// 		IsPublished:       in.IsPublished,
// 	}
// }

// func toModel(in Article) *models.Article {
// 	return &models.Article{
// 		ArticleURL:        in.ArticleURL,
// 		NewsArticleID:     in.NewsArticleID,
// 		PublishDate:       in.PublishDate,
// 		Taxonomies:        in.Taxonomies,
// 		TeaserText:        in.TeaserText,
// 		ThumbnailImageURL: in.ThumbnailImageURL,
// 		Title:             in.Title,
// 		OptaMatchId:       in.OptaMatchId,
// 		LastUpdateDate:    in.LastUpdateDate,
// 		IsPublished:       in.IsPublished,
// 	}
// }
