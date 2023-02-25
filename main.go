package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jasonlvhit/gocron"
	"github.com/sirupsen/logrus"
	"github.com/xavimg/articles/internal/models"
	"github.com/xavimg/articles/internal/services/articles"
)

var (
	listenAddr = ":4007"
	dbName     = "articles"
	mongoURL   = "mongodb://mongo:27017"
)

func main() {
	ctx := context.Background()
	if err := models.ConnectRepo(ctx); err != nil {
		log.Panic(err)
	}

	go func() {
		s := gocron.NewScheduler()
		s.Every(3).Seconds().Do(articles.PollingNews)
		<-s.Start()
	}()

	router := chi.NewRouter()
	articles.NewServer(router)

	logrus.Fatal(http.ListenAndServe(":4007", router))
}
