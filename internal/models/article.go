package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
