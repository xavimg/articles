package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jasonlvhit/gocron"
	"github.com/sirupsen/logrus"

	"github.com/xavimg/articles/internal/config"
	articles_controller "github.com/xavimg/articles/internal/controllers/articles"
	"github.com/xavimg/articles/internal/models"
	articles_service "github.com/xavimg/articles/internal/services/articles"
)

func main() {
	if err := config.LoadSettings(); err != nil {
		logrus.Fatal(err)
	}

	if err := models.ConnectRepo(context.Background()); err != nil {
		logrus.Fatal(err)
	}

	s := gocron.NewScheduler()
	if err := s.Every(12).Seconds().Do(articles_service.PollingNews); err != nil {
		logrus.Fatal(err)
	}
	s.Start()

	router := chi.NewRouter()
	articles_controller.NewServer(router)
	logrus.Fatal(http.ListenAndServe(":4007", router))
}
