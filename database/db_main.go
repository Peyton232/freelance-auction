package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client   *mongo.Client
	users    *mongo.Collection
	auctions *mongo.Collection
}

func Connect(mongoUri string, mongoDb string) *DB {

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Print(err)
		log.Print("\nDB connection failed in database package")
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client.Connect(ctx)
	return &DB{
		users:    client.Database(mongoDb).Collection("users"),
		auctions: client.Database(mongoDb).Collection("auctions"),
		client:   client,
	}
}

// Disconnect to MongoDB - could use, on further inspection looks like mongo is good at closing itself
// func (db *DB) closeDB() {
// 	err := db.client.Disconnect(context.Background())
// 	if err != nil {
// 		log.Print(err)
//      log.Print("\nunable to disconnect from DB in database package\n")
//      return nil
// 	}
// }
