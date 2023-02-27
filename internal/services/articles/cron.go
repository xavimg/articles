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

	// Using same variables as the example of Google Drive.
	count  = "50"
	teamId = "t94"
)

// PollingNews
//  1. Fetch XML data from getnewlistinformation.
//  2. Transform that data XML data into Golang model structure for MongoDB. Create array of strings externalIds, and insert the id articl'es from data provider.
//  3. List the existing articles from MongoDB passing by parameter teamId and externalIds
//  4. Once we have the existingArticles we can play with updates, for update articles checking LastUpdate; and creates, for create new ones when don't exists.
//  5. We range the array of articles we parsed from XML to our model Golang structure, and for each article try to find by NewsArticleID. If the article is found,
//     and LastUpdateDate is the same, it means we dont want update or insert. When the article is not found, we call getnewsarticleinformation, for insert the
//     more detailed info such us content or galleryImgURL.
//  6. If article exists, we will insert into the array of updates, if dont exist we will insert into array of creates.
//  7. Finally, check the lenght of each array, and do the needed operation that can be ReplaceOne article, or InsertMany articles.
func PollingNews() {
	logrus.Info("New polling at %v", time.Now())

	ctx := context.Background()

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

	if len(updates) > 0 {
		if err := models.Repo.ReplaceOne(ctx, updates); err != nil {
			logrus.Error(errors.Trace(err))
			return
		}
	}
	if len(creates) > 0 {
		if err := models.Repo.InsertMany(ctx, creates); err != nil {
			logrus.Error(errors.Trace(err))
			return
		}
	}

	return
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
		logrus.Error(errors.Trace(err))
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

func findArticle(articles []*models.Article, article *models.Article) *models.Article {
	for _, a := range articles {
		if a.NewsArticleID == article.NewsArticleID {
			return a
		}
	}

	return nil
}
