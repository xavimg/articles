package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jasonlvhit/gocron"
	"github.com/sirupsen/logrus"
	"github.com/xavimg/articles/internal/config"
	"github.com/xavimg/articles/internal/models"
	"github.com/xavimg/articles/internal/services/articles"
)

func main() {
	if err := config.LoadSettings(); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	if err := models.ConnectRepo(ctx); err != nil {
		log.Fatal(err)
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
