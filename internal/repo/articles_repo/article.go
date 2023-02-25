package articles_repo

import (
	"context"

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

func (as *ArticleRepository) InsertMany(ctx context.Context, articles []*models.Article) error {
	var articlesToInsert []interface{}
	articlesToInsert = append(articlesToInsert, articles)

	if _, err := models.Repo.Articles().InsertMany(ctx, articlesToInsert); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (as *ArticleRepository) InsertOne(ctx context.Context, article *models.Article) error {
	if _, err := models.Repo.Articles().InsertOne(ctx, article); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (as *ArticleRepository) GetAll(ctx context.Context) ([]models.Article, error) {
	cursor, err := models.Repo.Articles().Find(ctx, map[string]any{})
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var articles []models.Article
	for cursor.Next(context.Background()) {
		var article models.Article
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

func (as *ArticleRepository) GetByID(ctx context.Context, id string) (*models.Article, error) {
	var article *models.Article
	if err := models.Repo.Articles().FindOne(ctx, bson.M{"_id": id}).Decode(&article); err != nil {
		logrus.Error(err)
		return nil, err
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
