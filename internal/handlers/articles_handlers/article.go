package articles_handlers

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"github.com/xavimg/articles/internal/models"
	"github.com/xavimg/articles/internal/repo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	host     = "www.wearehullcity.co.uk"
	endpoint = "/api/incrowd/getnewlistinformation"
)

type Server struct {
	listenAddr string
	repo       repo.Repo
}

func NewServer(addr string, aRepo repo.Repo) *Server {
	return &Server{
		listenAddr: addr,
		repo:       aRepo,
	}
}

func (s *Server) Run() {
	router := chi.NewRouter()
	router.Get("/articles/{count}", makeHTTPHandleFunc(s.All))
	router.Get("/article/{id}", makeHTTPHandleFunc(s.ByID))

	http.ListenAndServe(s.listenAddr, router)
}

func (ah *Server) ByID(w http.ResponseWriter, r *http.Request) error {
	count := chi.URLParam(r, "count")

	// Show which come from feed provider.
	var articles []Article
	providerArticles, err := providerArticles(count)
	if err != nil {
		repoArticles, err := ah.repo.GetAll(context.Background())
		if err != nil {
			return err
		}
		for _, article := range repoArticles {
			articles = append(articles, Article{
				ArticleURL:    article.ArticleURL,
				NewsArticleID: article.NewsArticleID,
			})
		}
	}

	// ah.store.Insert()

	providerArticles.Status = "Succes"
	providerArticles.Data = articles
	// response := {
	// 	Status: "succes",
	// 	Data:   articles,
	// }
	return writeJSON(w, http.StatusOK, providerArticles)
}

func (ah *Server) All(w http.ResponseWriter, r *http.Request) error {
	count := chi.URLParam(r, "count")

	// Fetch coming from provider-feed.
	providerArticles, err := providerArticles(count)
	if err != nil {
		// Alternative fetch if our provider-feed fails.
		savedArticles, err := ah.repo.GetAll(context.Background())
		if err != nil {
			logrus.Printf("err %s\n", err)
			return err
		}
		logrus.Println(err)
		return writeJSON(w, http.StatusOK, savedArticles)
	}

	insertArticles := []models.Article{}
	for _, article := range providerArticles.Data {
		insertArticles = append(insertArticles, models.Article{
			ID:                primitive.NewObjectID(),
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
	}

	go func() error {
		if err := ah.repo.InsertMany(context.Background(), insertArticles); err != nil {
			logrus.Println(err)
			return err
		}
		return nil
	}()

	return writeJSON(w, http.StatusOK, providerArticles)
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func providerArticles(count string) (*ArticlesJSON, error) {
	params := url.Values{}
	params.Add("count", count)
	url := &url.URL{
		Scheme:   "https",
		Host:     host,
		Path:     endpoint,
		RawQuery: params.Encode(),
	}
	resp, err := http.Get(url.String())
	if err != nil {
		logrus.Printf("Error %s", err)
		return nil, err
	}

	respByte, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Printf("Error %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	var dataClean models.DataXMLALL
	if err := xml.Unmarshal(respByte, &dataClean); err != nil {
		logrus.Printf("Error %s", err)
		return nil, err
	}

	var article Article
	var articles []Article
	for _, item := range dataClean.NewsletterNewsItems.NewsletterNewsItem {
		article = Article{
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

	responsejson := &ArticlesJSON{
		Status: "succes",
		Data:   articles,
	}

	return responsejson, nil
}

func callOneItemByID() *models.DataJSONOne {
	resp, err := http.Get("https://www.wearehullcity.co.uk/api/incrowd/getnewsarticleinformation?id=443426")
	if err != nil {
		logrus.Printf("Error %s", err)
		return nil
	}

	respByte, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Printf("Error %s", err)
		return nil
	}
	defer resp.Body.Close()

	var dataClean models.DataXMLOne
	if err := xml.Unmarshal(respByte, &dataClean); err != nil {
		logrus.Printf("Error %s", err)
		return nil
	}

	body := models.Article{
		ArticleURL:        dataClean.NewsletterNewsItem.ArticleURL,
		NewsArticleID:     dataClean.NewsletterNewsItem.NewsArticleID,
		PublishDate:       dataClean.NewsletterNewsItem.PublishDate,
		Taxonomies:        dataClean.NewsletterNewsItem.Taxonomies,
		TeaserText:        dataClean.NewsletterNewsItem.TeaserText,
		ThumbnailImageURL: dataClean.NewsletterNewsItem.ThumbnailImageURL,
		Title:             dataClean.NewsletterNewsItem.Title,
		OptaMatchId:       dataClean.NewsletterNewsItem.OptaMatchId,
		LastUpdateDate:    dataClean.NewsletterNewsItem.LastUpdateDate,
		IsPublished:       dataClean.NewsletterNewsItem.IsPublished,
	}
	responsejson := &models.DataJSONOne{
		Status: "succes",
		Data:   body,
	}

	return responsejson
}
