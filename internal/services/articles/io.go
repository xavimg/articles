package articles

import (
	"encoding/xml"
	"time"

	"github.com/xavimg/articles/internal/models"
)

type ArticlesJSON struct {
	Status   string    `json:"status"`
	Data     []Article `json:"data"`
	Metadata Metadata  `json:"metadata"`
}

func (aj *ArticlesJSON) Serialize(articles []*models.Article) *ArticlesJSON {
	articlesJSON := &ArticlesJSON{}
	for _, article := range articles {
		articlesJSON.Status = "succes"
		articlesJSON.Data = append(articlesJSON.Data, Article{
			ArticleURL:        article.ArticleURL,
			NewsArticleID:     article.NewsArticleID,
			PublishDate:       article.PublishDate,
			Taxonomies:        article.Taxonomies,
			TeaserText:        article.TeaserText,
			ThumbnailImageURL: article.ThumbnailImageURL,
			Title:             article.Title,
			OptaMatchId:       article.OptaMatchId,
			LastUpdateDate:    article.LastUpdateDate,
			IsPublished:       article.IsPublished,
		})
		articlesJSON.Metadata = Metadata{
			CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
			TotalItems: len(articles),
			Sort:       article.IsPublished,
		}
	}
	return articlesJSON
}

type ArticleJSON struct {
	Status   string   `json:"status"`
	Data     Article  `json:"data"`
	Metadata Metadata `json:"metadata"`
}

func (aj *ArticleJSON) Serialize(article *models.Article) *ArticleJSON {
	articleJSON := &ArticleJSON{
		Status: "succes",
		Data: Article{
			ArticleURL:        article.ArticleURL,
			PublishDate:       article.PublishDate,
			Taxonomies:        article.Taxonomies,
			TeaserText:        article.TeaserText,
			ThumbnailImageURL: article.ThumbnailImageURL,
			Title:             article.Title,
			OptaMatchId:       article.OptaMatchId,
			LastUpdateDate:    article.LastUpdateDate,
			IsPublished:       article.IsPublished,
		},
		Metadata: Metadata{
			CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
			TotalItems: 1,
			Sort:       article.IsPublished,
		},
	}
	return articleJSON
}

type Metadata struct {
	CreatedAt  string `json:"createdAt"`
	TotalItems int    `json:"totalItems"`
	Sort       string `json:"sort"`
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
