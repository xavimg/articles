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

	// This call is because for testing the test, and call our endpoint we will need data in mongoDB, and we don't
	// want to wait the timer of our cron-task(5min). This 5 min wait to insert, only happends at first time, because
	// cron will do every 5 min, and first time, will be after 5 min.
	go articles_service.PollingNews()

	s := gocron.NewScheduler()
	if err := s.Every(5).Minutes().Do(articles_service.PollingNews); err != nil {
		logrus.Fatal(err)
	}
	s.Start()

	router := chi.NewRouter()
	articles_controller.NewServer(router)
	logrus.Fatal(http.ListenAndServe(":4007", router))
}
