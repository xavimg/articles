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
	scheme       = "https"
	host         = "www.wearehullcity.co.uk"
	endpointList = "/api/incrowd/getnewlistinformation"
	endpointOne  = "/api/incrowd/getnewsarticleinformation"
)

func PollingNews() {
	logrus.Info("New polling at %v", time.Now())

	params := url.Values{}
	params.Add("count", "50") // Im using 50 es the example in drive.
	url := &url.URL{
		Scheme:   scheme,
		Host:     host,
		Path:     endpointList,
		RawQuery: params.Encode(),
	}

	res, err := http.Get(url.String())
	if err != nil {
		logrus.Errorf("error %s", err)
		return
	}

	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		logrus.Errorf("eror %s", err)
		return
	}
	defer res.Body.Close()

	var providerData DataXML
	if err := xml.Unmarshal(resBytes, &providerData); err != nil {
		logrus.Errorf("error %s", err)
		return
	}

	articles := DeserializeXML(providerData)

	go func() error {
		if err := models.Repo.InsertManyTask(context.Background(), articles); err != nil {
			logrus.Errorf("error %s", err)
			return err
		}
		return nil
	}()

	return
}

func DeserializeXML(providerData DataXML) []*models.Article {
	var article *models.Article
	var articles []*models.Article

	for _, xml := range providerData.NewsletterNewsItems.NewsletterNewsItem {
		article = &models.Article{
			ArticleURL:        xml.ArticleURL,
			NewsArticleID:     xml.NewsArticleID,
			PublishDate:       xml.PublishDate,
			Taxonomies:        xml.Taxonomies,
			TeaserText:        xml.TeaserText,
			ThumbnailImageURL: xml.ThumbnailImageURL,
			Title:             xml.Title,
			OptaMatchId:       xml.OptaMatchId,
			LastUpdateDate:    xml.LastUpdateDate,
			IsPublished:       xml.IsPublished,
		}
		articles = append(articles, article)
	}

	return articles
}
