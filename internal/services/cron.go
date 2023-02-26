package services

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/xavimg/articles/internal/config"
	"github.com/xavimg/articles/internal/dtos"
	"github.com/xavimg/articles/internal/models"
)

const (
	endpointList = "getnewlistinformation"

	count = "50"

	// Refference to Hull City teamId due to api requirements as test.
	teamId = "t94"
)

func PollingNews() {
	logrus.Info("New polling at %v", time.Now())

	url := fmt.Sprintf("%s/%s?count=%s", config.Settings.FeedProvider.HullCity, endpointList, count)
	res, err := http.Get(url)
	if err != nil {
		logrus.Errorf("error %s", err)
		return
	}
	if res.StatusCode != http.StatusOK {
		logrus.Errorf("error %s", err)
		return
	}

	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		logrus.Errorf("eror %s", err)
		return
	}
	defer res.Body.Close()

	var providerData dtos.ArticleXML
	if err := xml.Unmarshal(resBytes, &providerData); err != nil {
		logrus.Errorf("error %s", err)
		return
	}

	var articles []*models.Article
	for _, xml := range providerData.NewsletterNewsItems.NewsletterNewsItem {
		date, err := time.Parse("2006-01-02 15:04:05", xml.LastUpdates)
		if err != nil {
			logrus.Error(err)
			return
		}
		article := &models.Article{
			TeamID:            teamId,
			ArticleURL:        xml.ArticleURL,
			NewsArticleID:     xml.NewsArticleID,
			PublishDate:       xml.PublishDate,
			Type:              xml.Taxonomies,
			TeaserText:        xml.TeaserText,
			ThumbnailImageURL: xml.ThumbnailImageURL,
			Title:             xml.Title,
			LastUpdateDate:    date,
			OptaMatchId:       xml.OptaMatchId,
			IsPublished:       xml.IsPublished,
		}

		articles = append(articles, article)
	}

	fmt.Println("not inserted")

	if err := models.Repo.BatchCreate(context.Background(), articles); err != nil {
		logrus.Error(err)
		return
	}

	fmt.Println("inserted")

	return
}
