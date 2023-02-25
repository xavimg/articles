package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
	"github.com/xavimg/articles/internal/handlers/articles_handlers"
	"github.com/xavimg/articles/internal/repo/articles_repo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	listenAddr = ":4007"
	dbName     = "articles"
	mongoURL   = "mongodb://mongo:27017"
)

func main() {
	// connect to mongo
	ctx := context.TODO()
	mongoClient, err := connectToMongo(ctx)
	if err != nil {
		log.Panic(err)
	}

	articlesRepo := articles_repo.NewRepository(mongoClient)

	server := articles_handlers.NewServer(listenAddr, articlesRepo)

	fmt.Println("running on port ", listenAddr)
	server.Run()
}

func connectToMongo(ctx context.Context) (*mongo.Database, error) {
	// create connection options
	connection := options.Client().ApplyURI(mongoURL)

	// connect
	c, err := mongo.Connect(context.TODO(), connection)
	if err != nil {
		logrus.Println("Error connecting:", err)
		return nil, err
	}

	logrus.Println("Connected to mongo!")

	return c.Database("articles"), nil
}
