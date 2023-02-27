package articles

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/juju/errors"
	"github.com/sirupsen/logrus"
	"github.com/xavimg/articles/internal/config"
	"github.com/xavimg/articles/internal/dtos"
	"github.com/xavimg/articles/internal/models"
)

const (
	endpointList     = "getnewlistinformation"
	endpointDetailed = "getnewsarticleinformation"

	count = "50"

	// We don't know the other teamId we should poll. So as example, I will use the t94 as refference, like the example.
	teamId = "t94"
)

func PollingNews() {
	timeStart := time.Now()
	ctx := context.Background()
	logrus.Info("New polling at %v", time.Now())

	providerData, err := listArticles()
	if err != nil {
		logrus.Error(errors.Trace(err))
		return
	}

	externalIds := []string{}
	articles := []*models.Article{}
	for _, news := range providerData.NewsletterNewsItems.NewsletterNewsItem {
		externalIds = append(externalIds, news.NewsArticleID)
		articles = append(articles, &models.Article{
			TeamID:            teamId,
			ArticleURL:        news.ArticleURL,
			NewsArticleID:     news.NewsArticleID,
			PublishDate:       news.ParsePublishDate(),
			Type:              news.Taxonomies,
			TeaserText:        news.TeaserText,
			ThumbnailImageURL: news.ThumbnailImageURL,
			Title:             news.Title,
			LastUpdateDate:    news.ParseLastUpdate(),
			OptaMatchId:       news.OptaMatchId,
			IsPublished:       news.IsPublished,
		})
	}

	existingArticles, err := models.Repo.ListArticlesByIds(ctx, teamId, externalIds)
	if err != nil {
		logrus.Error(errors.Trace(err))
		return
	}

	updates := []*models.Article{}
	creates := []interface{}{}
	for _, article := range articles {
		existing := findArticle(existingArticles, article)
		if existing != nil && existing.LastUpdateDate.Equal(article.LastUpdateDate) {
			continue
		}

		detail, err := getArticle(article.NewsArticleID)
		if err != nil {
			logrus.Error(errors.Trace(err))
			return
		}

		article.Subtitle = detail.NewsletterNewsItems.Subtitle
		article.GalleryURLs = detail.NewsletterNewsItems.GalleryImageURLs
		article.VideoURL = detail.NewsletterNewsItems.VideoURL
		article.Content = detail.NewsletterNewsItems.Content

		if existing != nil {
			updates = append(updates, article)
			continue
		}

		creates = append(creates, article)
	}

	elapsed := time.Since(timeStart)
	fmt.Println("finish", elapsed)

	if len(updates) > 0 {
		models.Repo.ReplaceOne(ctx, updates)
		elapsed := time.Since(timeStart)
		fmt.Println("updated", elapsed)
	}
	if len(creates) > 0 {
		models.Repo.InsertMany(ctx, creates)
		elapsed := time.Since(timeStart)
		fmt.Println("insert", elapsed)
	}

	return
}

func findArticle(articles []*models.Article, article *models.Article) *models.Article {
	for _, a := range articles {
		if a.NewsArticleID == article.NewsArticleID {
			return a
		}
	}

	return nil
}

func listArticles() (*dtos.NewListInformation, error) {
	url := fmt.Sprintf("%s/%s?count=%s", config.Settings.FeedProvider.HullCity, endpointList, count)
	providerData := &dtos.NewListInformation{}

	if err := call(url, providerData); err != nil {
		return nil, err
	}

	return providerData, nil
}

func getArticle(id string) (*dtos.NewsArticleInformation, error) {
	url := fmt.Sprintf("%s/%s?id=%s", config.Settings.FeedProvider.HullCity, endpointDetailed, id)
	providerData := &dtos.NewsArticleInformation{}

	if err := call(url, providerData); err != nil {
		return nil, err
	}

	return providerData, nil
}

func call(url string, reply interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		logrus.Error(errors.Trace(err))
		return err
	}
	if res.StatusCode != http.StatusOK {
		logrus.Error(errors.Trace(err))
		return err
	}

	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		logrus.Error(errors.Trace(err))
		return err
	}
	defer res.Body.Close()

	if err := xml.Unmarshal([]byte(resBytes), &reply); err != nil {
		logrus.Error(errors.Trace(err))
		return err
	}

	return nil
}
