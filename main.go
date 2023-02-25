package main

import (
	"context"
	"fmt"
	"log"
	"time"

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

var client *mongo.Client

func main() {
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}
	client = mongoClient

	// create a context in order to disconnect
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// close connection
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	articlesRepo := articles_repo.NewRepository(client.Database(dbName))
	server := articles_handlers.NewServer(listenAddr, articlesRepo)
	fmt.Println("running on port ", listenAddr)
	server.Run()

}

// func connectMongoDB() *mongo.Client {
// 	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
// 	if err != nil {
// 		logrus.Fatal(err)
// 	}
// 	if err := client.Connect(context.TODO()); err != nil {
// 		logrus.Fatal(err)
// 	}

// 	return client
// }

func connectToMongo() (*mongo.Client, error) {
	// create connection options
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	// connect
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error connecting:", err)
		return nil, err
	}

	log.Println("Connected to mongo!")

	return c, nil
}
