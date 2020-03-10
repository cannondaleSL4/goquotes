package mongodbService

import (
	"context"
	"fmt"
	tinkoff "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"github.com/goquotes/constants"
	"go.mongodb.org/mongo-driver/bson"
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

func InsertNewQuotes(client *Client, stock [][]tinkoff.Candle) {
	collection := client.Database(constants.DBNAME).Collection("stocks")

	insertStocks := []interface{}{}
	for _, t := range stock {
		insertStocks = append(insertStocks, t)
	}

	collection.InsertMany(context.TODO(), insertStocks, options.InsertMany().SetOrdered(false))
}

func getQuotes(client *Client, figi string, num_limit int64) {
	collection := client.Database(constants.DBNAME).Collection("stocks")

	findOptions := options.Find()
	findOptions.SetLimit(num_limit)
	findOptions.SetSort(map[string]int{"ts": 1})

	filter := bson.M{
		"ts": bson.M{
			"$gte": time.Now().AddDate(0, 0, int(-num_limit)).UTC(),
		},
		"figi": figi,
	}

	var results []*tinkoff.Candle

	cur, err := collection.Find(context.TODO(), filter, findOptions)

	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem tinkoff.Candle
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	fmt.Println()
}

func closeConnection() {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}
