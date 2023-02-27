package dtos

import (
	"encoding/xml"
	"time"

	"github.com/sirupsen/logrus"
)

type NewListInformation struct {
	XMLName             xml.Name            `xml:"NewListInformation"`
	ClubName            string              `xml:"ClubName"`
	ClubWebsiteURL      string              `xml:"ClubWebsiteURL"`
	NewsletterNewsItems NewsletterNewsItems `xml:"NewsletterNewsItems"`
}

type NewsletterNewsItems struct {
	XMLName            xml.Name             `xml:"NewsletterNewsItems"`
	NewsletterNewsItem []NewsletterNewsItem `xml:"NewsletterNewsItem"`
}

type NewsletterNewsItem struct {
	XMLName           xml.Name `xml:"NewsletterNewsItem"`
	ArticleURL        string   `xml:"ArticleURL"`
	NewsArticleID     string   `xml:"NewsArticleID"`
	PublishDate       string   `xml:"PublishDate"`
	Taxonomies        []string `xml:"Taxonomies"`
	TeaserText        string   `xml:"TeaserText"`
	ThumbnailImageURL string   `xml:"ThumbnailImageURL"`
	Title             string   `xml:"Title"`
	OptaMatchId       string   `xml:"OptaMatchId"`
	LastUpdates       string   `xml:"LastUpdateDate"`
	IsPublished       bool     `xml:"IsPublished"`
}

func (n *NewsletterNewsItem) ParsePublishDate() time.Time {
	date, err := time.Parse("2006-01-02 15:04:05", n.PublishDate)
	if err != nil {
		logrus.Error(err)
		return time.Time{}
	}
	return date
}

func (n *NewsletterNewsItem) ParseLastUpdate() time.Time {
	date, err := time.Parse("2006-01-02 15:04:05", n.LastUpdates)
	if err != nil {
		logrus.Error(err)
		return time.Time{}
	}
	return date
}

type NewsArticleInformation struct {
	XMLName             xml.Name    `xml:"NewsArticleInformation"`
	ClubName            string      `xml:"ClubName"`
	ClubWebsiteURL      string      `xml:"ClubWebsiteURL"`
	NewsletterNewsItems NewsArticle `xml:"NewsArticle"`
}

type NewsArticle struct {
	XMLName           xml.Name `xml:"NewsArticle"`
	ArticleURL        string   `xml:"ArticleURL"`
	NewsArticleID     string   `xml:"NewsArticleID"`
	PublishDate       string   `xml:"PublishDate"`
	Taxonomies        []string `xml:"Taxonomies"`
	TeaserText        string   `xml:"TeaserText"`
	Subtitle          string   `xml:"Subtitle"`
	ThumbnailImageURL string   `xml:"ThumbnailImageURL"`
	Content           string   `xml:"BodyText"`
	GalleryImageURLs  []string `xml:"GalleryImageURLs"`
	OptaMatchId       string   `xml:"OptaMatchId"`
	VideoURL          string   `xml:"VideoURL"`
	LastUpdates       string   `xml:"LastUpdateDate"`
	IsPublished       bool     `xml:"IsPublished"`
}
