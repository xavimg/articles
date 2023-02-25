package models

import (
	"encoding/xml"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DataXMLALL struct {
	XMLName             xml.Name `xml:"NewListInformation"`
	ClubName            string   `xml:"ClubName"`
	ClubWebsiteURL      string   `xml:"ClubWebsiteURL"`
	NewsletterNewsItems struct {
		NewsletterNewsItem []struct {
			ArticleURL        string `xml:"ArticleURL"`
			NewsArticleID     string `xml:"NewsArticleID"`
			PublishDate       string `xml:"PublishDate"`
			Taxonomies        string `xml:"Taxonomies"`
			TeaserText        string `xml:"TeaserText"`
			ThumbnailImageURL string `xml:"ThumbnailImageURL"`
			Title             string `xml:"Title"`
			OptaMatchId       string `xml:"OptaMatchId"`
			LastUpdateDate    string `xml:"LastUpdateDate"`
			IsPublished       string `xml:"IsPublished"`
		} `xml:"NewsletterNewsItem"`
	} `xml:"NewsletterNewsItems"`
}

type DataXMLOne struct {
	XMLName            xml.Name `xml:"NewsArticleInformation"`
	ClubName           string   `xml:"ClubName"`
	ClubWebsiteURL     string   `xml:"ClubWebsiteURL"`
	NewsletterNewsItem struct {
		ArticleURL        string `xml:"ArticleURL"`
		NewsArticleID     string `xml:"NewsArticleID"`
		PublishDate       string `xml:"PublishDate"`
		Taxonomies        string `xml:"Taxonomies"`
		TeaserText        string `xml:"TeaserText"`
		ThumbnailImageURL string `xml:"ThumbnailImageURL"`
		Title             string `xml:"Title"`
		OptaMatchId       string `xml:"OptaMatchId"`
		LastUpdateDate    string `xml:"LastUpdateDate"`
		IsPublished       string `xml:"IsPublished"`
	} `xml:"NewsArticle"`
}

type DataJSONAll struct {
	Status string    `json:"status"`
	Data   []BodyAll `json:"data"`
	// Metadata Metadata `json:"metadata"`
}

type DataJSONOne struct {
	Status string  `json:"status"`
	Data   Article `json:"data"`
	// Metadata Metadata `json:"metadata"`
}

type BodyAll struct {
	ArticleURL        string `json:"articleURL"`
	NewsArticleID     string `json:"newsArticleID"`
	PublishDate       string `json:"publishDate"`
	Taxonomies        string `json:"taxonomies"`
	TeaserText        string `json:"teaserText"`
	ThumbnailImageURL string `json:"thumbnailImageURL"`
	Title             string `json:"title"`
	OptaMatchId       string `json:"optaMatchID"`
	LastUpdateDate    string `json:"lastUpdateDate"`
	IsPublished       string `json:"published"`
}

// type Article struct {
// 	ID                primitive.ObjectID `bson:"_id,omitempty"`
// 	ArticleURL        string             `json:"articleURL"`
// 	NewsArticleID     string             `json:"newsArticleID"`
// 	PublishDate       string             `json:"publishDate"`
// 	Taxonomies        string             `json:"taxonomies"`
// 	TeaserText        string             `json:"teaserText"`
// 	ThumbnailImageURL string             `json:"thumbnailImageURL"`
// 	Title             string             `json:"title"`
// 	OptaMatchId       string             `json:"optaMatchID"`
// 	LastUpdateDate    string             `json:"lastUpdateDate"`
// 	IsPublished       string             `json:"published"`
// }

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
