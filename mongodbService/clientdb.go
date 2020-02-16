package mongodbService

import (
	"context"
	"flag"
	"fmt"
	quotes "github.com/goquotes/const"
	"log"
	"time"

	. "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *Client
	mongoURL = "mongodb://127.0.0.1:27017"
	cmd          = flag.String("cmd", "", "list or add?")
	address    = flag.String("address", "", "mongodb address to connect to")
	database   = flag.String("db", "", "The name of the database to connect to")
	collection = flag.String("collection", "", "The collection (in the db) to connect to")
	key        = flag.String("field", "", "The field you'd like to place an index on")
	unique     = flag.Bool("unique", false, "Would you like the index to be unique?")
	value      = flag.Int("type", 1, "would you like the index to be ascending (1) or descending (-1)?")
)

func GetClient() *Client {

	client, err := NewClient(options.Client().ApplyURI(mongoURL))


	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)


	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

func InsertNewQuotes( client *Client, stock [] quotes.StocksFromResponse) {
	collection := client.Database("quotes").Collection("stocks")

	insertStocks := []interface{}{}
	for _, t := range stock {
		insertStocks = append(insertStocks, t)
	}

	insetResult, err := collection.InsertMany(context.TODO(), insertStocks)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insetResult.InsertedIDs)
}

func closeConnection()  {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}

