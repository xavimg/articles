package store

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/xavimg/incrowd-test/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ArticleRepository struct {
	db *mongo.Database
}

func NewArticleRepository(db *mongo.Database) *ArticleRepository {
	return &ArticleRepository{db: db}
}

type Article struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	ArticleURL        string             `bson:"articleURL,omitempty"`
	NewsArticleID     string             `bson:"newsArticleID,omitempty"`
	PublishDate       string             `bson:"publishDate,omitempty"`
	Taxonomies        string             `bson:"taxonomies,omitempty"`
	TeaserText        string             `bson:"teaserText,omitempty"`
	ThumbnailImageURL string             `bson:"thumbnailImageURL,omitempty"`
	Title             string             `bson:"title,omitempty"`
	OptaMatchId       string             `bson:"optaMatchID,omitempty"`
	LastUpdateDate    string             `bson:"lastUpdateDate,omitempty"`
	IsPublished       string             `bson:"published,omitempty"`
}

func (as *ArticleRepository) Insert(ctx context.Context, article *models.Article) error {
	_, err := as.db.Collection("articles").InsertOne(ctx, fromModel(article))
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

func fromModel(in *models.Article) *Article {
	return &Article{
		ArticleURL:        in.ArticleURL,
		NewsArticleID:     in.NewsArticleID,
		PublishDate:       in.PublishDate,
		Taxonomies:        in.Taxonomies,
		TeaserText:        in.TeaserText,
		ThumbnailImageURL: in.ThumbnailImageURL,
		Title:             in.Title,
		OptaMatchId:       in.OptaMatchId,
		LastUpdateDate:    in.LastUpdateDate,
		IsPublished:       in.IsPublished,
	}
}

func toModel(in Article) *models.Article {
	return &models.Article{
		ArticleURL:        in.ArticleURL,
		NewsArticleID:     in.NewsArticleID,
		PublishDate:       in.PublishDate,
		Taxonomies:        in.Taxonomies,
		TeaserText:        in.TeaserText,
		ThumbnailImageURL: in.ThumbnailImageURL,
		Title:             in.Title,
		OptaMatchId:       in.OptaMatchId,
		LastUpdateDate:    in.LastUpdateDate,
		IsPublished:       in.IsPublished,
	}
}
