package dtos

import (
	"encoding/xml"
	"time"
)

type ListReply struct {
	Status   string     `json:"status"`
	Data     []*Article `json:"data"`
	Metadata *Metadata  `json:"metadata"`
}

type Metadata struct {
	CreatedAt  string `json:"createdAt"`
	TotalItems int    `json:"totalItems"`
	Sort       string `json:"sort"`
}

type Article struct {
	ID                string    `json:"id"`
	TeamID            string    `json:"teamId"`
	ArticleURL        string    `json:"articleURL"`
	NewsArticleID     string    `json:"newsArticleID"`
	PublishDate       string    `json:"publishDate"`
	Type              []string  `json:"type"`
	TeaserText        string    `json:"teaserText"`
	ThumbnailImageURL string    `json:"thumbnailImageURL"`
	Title             string    `json:"title"`
	OptaMatchId       string    `json:"optaMatchID"`
	LastUpdateDate    time.Time `json:"lastUpdateDate"`
	IsPublished       bool      `json:"published"`
}

type ArticleXML struct {
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
