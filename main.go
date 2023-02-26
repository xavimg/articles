package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jasonlvhit/gocron"
	"github.com/sirupsen/logrus"
	"github.com/xavimg/articles/internal/config"
	"github.com/xavimg/articles/internal/controllers"
	"github.com/xavimg/articles/internal/models"
	"github.com/xavimg/articles/internal/services"
)

func main() {
	if err := config.LoadSettings(); err != nil {
		logrus.Fatal(err)
	}

	if err := models.ConnectRepo(context.Background()); err != nil {
		logrus.Fatal(err)
	}

	s := gocron.NewScheduler()
	s.Every(3).Seconds().Do(services.PollingNews)
	s.Start()

	router := chi.NewRouter()
	controllers.NewServer(router)

	logrus.Fatal(http.ListenAndServe(":4007", router))
}
