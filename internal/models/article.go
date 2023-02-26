package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	TeamID            string             `bson:"teamId,omitempty"`
	ArticleURL        string             `bson:"articleURL,omitempty"`
	NewsArticleID     string             `bson:"newsArticleID,omitempty"`
	PublishDate       string             `bson:"publishDate,omitempty"`
	Type              []string           `bson:"taxonomies,omitempty"`
	TeaserText        string             `bson:"teaserText,omitempty"`
	ThumbnailImageURL string             `bson:"thumbnailImageURL,omitempty"`
	Title             string             `bson:"title,omitempty"`
	OptaMatchId       string             `bson:"optaMatchID,omitempty"`
	LastUpdateDate    time.Time          `bson:"lastUpdateDate,omitempty"`
	IsPublished       bool               `bson:"published,omitempty"`
}
