package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
} 
func Connect() *DB{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoURL := os.Getenv("MONGO_URL")
	if mongoURL==""{
		log.Fatal("Error laoding Mongo url from env ")
	}
	
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err!=nil{
		log.Fatalln(err.Error() + "This is the error for mongo db connect")
	}
	return &DB{
		client: client,
	}
}
