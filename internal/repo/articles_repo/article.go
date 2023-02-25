package articles_repo

import (
	"context"
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
	"github.com/xavimg/articles/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ArticleRepository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (as *ArticleRepository) InsertMany(ctx context.Context, articles []models.Article) error {
	var articlesToInsert []interface{}
	for _, article := range articles {
		articlesToInsert = append(articlesToInsert, article)
	}

	collection := as.db.Collection("articles")

	_, err := collection.InsertMany(ctx, articlesToInsert)
	if err != nil {
		logrus.Println(err)
	}
	return nil
}

func (as *ArticleRepository) InsertOne(ctx context.Context, article *models.Article) error {
	collection := as.db.Collection("articles")

	_, err := collection.InsertOne(ctx, article)
	if err != nil {
		logrus.Println(err)
	}
	return nil
}

func (as *ArticleRepository) GetAll(ctx context.Context) ([]models.Article, error) {
	cursor, err := as.db.Collection("articles").Find(ctx, map[string]any{})
	if err != nil {
		logrus.Println("error: ", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var articles []models.Article
	for cursor.Next(context.Background()) {
		var article models.Article
		if err := cursor.Decode(&article); err != nil {
			log.Fatal(err)
		}
		articles = append(articles, article)
		fmt.Println(article)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return articles, nil
}

func (as *ArticleRepository) GetByID(ctx context.Context, id string) (*models.Article, error) {
	var article *models.Article
	if err := as.db.Collection("articles").FindOne(ctx, bson.M{"_id": id}).Decode(&article); err != nil {
		log.Fatal(err)
	}
	return article, nil
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
