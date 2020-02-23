package mongodbService

import (
	"context"
	"fmt"
	tinkoff "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"github.com/goquotes/constants"
	"log"
	"strings"
	"time"

	. "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client   *Client
	mongoURL = "mongodb://127.0.0.1:27017"
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

func InsertConst(client *Client) {
	djDbNameLower := strings.ToLower(strings.Replace(constants.DOWJONES, " ", "", -1))

	collection := client.Database(constants.DBNAME).Collection(djDbNameLower)
	collection.Drop(context.TODO())

	for _, l := range constants.QuotesMapDJ {
		_, err := collection.InsertOne(context.TODO(), l)
		if err != nil {
			log.Fatal(err)
		}
	}

	rusDbNameLower := strings.ToLower(strings.Replace(constants.RUS, " ", "", -1))

	collectionRus := client.Database(constants.DBNAME).Collection(rusDbNameLower)
	collectionRus.Drop(context.TODO())

	for _, l := range constants.QuotesMapRUS {
		_, err := collectionRus.InsertOne(context.TODO(), l)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func InsertNewQuotes(client *Client, stock []tinkoff.Candle) {
	collection := client.Database(constants.DBNAME).Collection("stocks")

	insertStocks := []interface{}{}
	for _, t := range stock {
		insertStocks = append(insertStocks, t)
	}

	_, err := collection.InsertMany(context.TODO(), insertStocks)
	if err != nil {
		log.Fatal(err)
	}
}

func closeConnection() {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}
