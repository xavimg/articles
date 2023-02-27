package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	NewsArticleID     string             `bson:"newsArticleID,omitempty"`
	TeamID            string             `bson:"teamId,omitempty"`
	OptaMatchId       string             `bson:"optaMatchID,omitempty"`
	Title             string             `bson:"title,omitempty"`
	Subtitle          string             `bson:"subtitle,omitempty"`
	Type              []string           `bson:"type,omitempty"`
	TeaserText        string             `bson:"teaser,omitempty"`
	Content           string             `bson:"content,omitempty"`
	ArticleURL        string             `bson:"url,omitempty"`
	ThumbnailImageURL string             `bson:"imageUrl,omitempty"`
	GalleryURLs       []string           `bson:"galleryUrls,omitempty"`
	VideoURL          string             `bson:"videoUrl,omitempty"`
	PublishDate       time.Time          `bson:"published,omitempty"`
	LastUpdateDate    time.Time          `bson:"lastUpdateDate,omitempty"`
	IsPublished       bool               `bson:"isPublished,omitempty"`
}
