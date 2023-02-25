package articles_repo

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/xavimg/articles/internal/models"
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

	_, err := as.db.Collection("articles").InsertMany(ctx, articlesToInsert)
	if err != nil {
		logrus.Println(err)
	}
	return nil
}

func (as *ArticleRepository) GetAll(ctx context.Context) ([]*models.Article, error) {
	cursor, err := as.db.Collection("articles").Find(ctx, map[string]any{})
	if err != nil {
		return nil, err
	}

	articles := []*models.Article{}
	if err := cursor.All(ctx, articles); err != nil {
		return nil, err
	}

	return articles, nil
}

func (as *ArticleRepository) GetByID(ctx context.Context, id string) (*models.Article, error) {
	return nil, nil
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
