package articles

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
	"github.com/xavimg/articles/internal/models"
)

func PollingNews() {
	params := url.Values{}
	params.Add("count", "2")
	url := &url.URL{
		Scheme:   "https",
		Host:     host,
		Path:     endpointList,
		RawQuery: params.Encode(),
	}
	resp, err := http.Get(url.String())
	if err != nil {
		logrus.Printf("Error %s", err)
		return
	}

	respByte, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Printf("Error %s", err)
		return
	}
	defer resp.Body.Close()

	var dataClean models.DataXMLALL
	if err := xml.Unmarshal(respByte, &dataClean); err != nil {
		logrus.Printf("Error %s", err)
		return
	}

	var article *models.Article
	var articles []*models.Article
	for _, item := range dataClean.NewsletterNewsItems.NewsletterNewsItem {
		article = &models.Article{
			ArticleURL:        item.ArticleURL,
			NewsArticleID:     item.NewsArticleID,
			PublishDate:       item.PublishDate,
			Taxonomies:        item.Taxonomies,
			TeaserText:        item.TeaserText,
			ThumbnailImageURL: item.ThumbnailImageURL,
			Title:             item.Title,
			OptaMatchId:       item.OptaMatchId,
			LastUpdateDate:    item.LastUpdateDate,
			IsPublished:       item.IsPublished,
		}

		articles = append(articles, article)
	}

	go func() error {
		if err := models.Repo.InsertManyTask(context.Background(), articles); err != nil {
			logrus.Println(err)
			return err
		}
		return nil
	}()

	return
}
