package articles

import (
	"encoding/xml"
	"time"
)

type ArticlesJSON struct {
	Status   string    `json:"status"`
	Data     []Article `json:"data"`
	Metadata Metadata  `json:"metadata"`
}

type ArticleJSON struct {
	Status   string   `json:"status"`
	Data     Article  `json:"data"`
	Metadata Metadata `json:"metadata"`
}

type Metadata struct {
	CreatedAt  time.Time `json:"createdAt"`
	TotalItems int       `json:"totalItems"`
	Sort       string    `json:"sort"`
}

type Article struct {
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

type DataXML struct {
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
