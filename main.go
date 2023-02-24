package main

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"

	"github.com/xavimg/incrowd-test/internal/handlers"

	"github.com/sirupsen/logrus"
	"github.com/xavimg/incrowd-test/internal/store"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DataXMLALL struct {
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

type DataXMLOne struct {
	XMLName            xml.Name `xml:"NewsArticleInformation"`
	ClubName           string   `xml:"ClubName"`
	ClubWebsiteURL     string   `xml:"ClubWebsiteURL"`
	NewsletterNewsItem struct {
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
	} `xml:"NewsArticle"`
}

type DataJSONAll struct {
	Status string    `json:"status"`
	Data   []BodyAll `json:"data"`
	// Metadata Metadata `json:"metadata"`
}

type DataJSONOne struct {
	Status string  `json:"status"`
	Data   BodyOne `json:"data"`
	// Metadata Metadata `json:"metadata"`
}

type BodyAll struct {
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

type BodyOne struct {
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

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		logrus.Fatal(err)
	}
	if err := client.Connect(context.TODO()); err != nil {
		logrus.Fatal(err)
	}

	articleRepository := store.NewArticleRepository(client.Database("articles"))

	server := handlers.NewArticleHandler(":4007", articleRepository)

	server.Run()

	// r := chi.NewRouter()

	// r.Get("/all-items", articleHandlers. articleHandlers.AllArticles)

	// r.Get("/item", func(w http.ResponseWriter, r *http.Request) {
	// 	res := callOneItemByID()
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Header().Add("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(res)
	// })

}

func callAllItems() *DataJSONAll {
	resp, err := http.Get("https://www.wearehullcity.co.uk/api/incrowd/getnewlistinformation?count=50")
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

	var dataClean DataXMLALL
	if err := xml.Unmarshal(respByte, &dataClean); err != nil {
		logrus.Printf("Error %s", err)
		return nil
	}

	var body BodyAll
	var bodys []BodyAll
	for _, item := range dataClean.NewsletterNewsItems.NewsletterNewsItem {
		body = BodyAll{
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

		bodys = append(bodys, body)
	}
	responsejson := &DataJSONAll{
		Status: "succes",
	}
	responsejson.Data = bodys

	return responsejson
}

func callOneItemByID() *DataJSONOne {
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

	var dataClean DataXMLOne
	if err := xml.Unmarshal(respByte, &dataClean); err != nil {
		logrus.Printf("Error %s", err)
		return nil
	}

	body := BodyOne{
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
	responsejson := &DataJSONOne{
		Status: "succes",
		Data:   body,
	}

	return responsejson
}
