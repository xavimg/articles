package articles

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/xavimg/articles/internal/models"
)

const (
	host         = "www.wearehullcity.co.uk"
	endpointList = "/api/incrowd/getnewlistinformation"
	endpointOne  = "/api/incrowd/getnewsarticleinformation"
)

func PollingNews() {
	logrus.Info("New polling at %v", time.Now())

	params := url.Values{}
	params.Add("count", "5") // He de hardcoldearlo porque nose como hacero dinámico
	url := &url.URL{
		Scheme:   "https",
		Host:     host,
		Path:     endpointList,
		RawQuery: params.Encode(),
	}

	res, err := http.Get(url.String())
	if err != nil {
		logrus.Printf("Error %s", err)
		return
	}

	resByte, err := io.ReadAll(res.Body)
	if err != nil {
		logrus.Printf("Error %s", err)
		return
	}
	defer res.Body.Close()

	var providerData DataXML
	if err := xml.Unmarshal(resByte, &providerData); err != nil {
		logrus.Printf("Error %s", err)
		return
	}

	var article *models.Article
	var articles []*models.Article
	for _, item := range providerData.NewsletterNewsItems.NewsletterNewsItem {
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
