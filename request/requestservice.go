package request

import (
	"context"
	"flag"
	"log"

	//"fmt"
	quotes "github.com/goquotes/constants"
	//"log"
	"os"
	"time"

	tinkoff "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
)

var token = flag.String("token", os.Getenv("TOKEN"), "your token")

func UpdateFromTo(from time.Time, to time.Time) []tinkoff.Candle {
	var arrayOfRequestData []quotes.RequestData
	for _, element := range quotes.GetQuotesDJ() {
		var req quotes.RequestData
		req.FIGI = element
		req.From = from
		req.To = to
		req.Resolution = tinkoff.CandleInterval1Hour
		arrayOfRequestData = append(arrayOfRequestData, req)
	}
	resp := requestToServer(arrayOfRequestData)
	if resp != nil && len(resp) != 0 {
		return resp
	}
	return nil
}

//func requestToServer(arrayOfRequestData []quotes.RequestData) []quotes.StocksFromResponse {
func requestToServer(arrayOfRequestData []quotes.RequestData) []tinkoff.Candle {
	//stocks := make([]quotes.StocksFromResponse, 0)
	stocks := make([]tinkoff.Candle, 0)

	for _, data := range arrayOfRequestData {
		candles, err := makeRequest(data)
		if err != nil {
			log.Printf("%+v\n", err)
			return nil
		}

		if len(candles) == 0 {
			log.Printf("Len of the reponse array is 0 for data: %+v\n", data.From)
			return nil
		}
		stocks = append(stocks, candles...)
	}

	return stocks
}

func makeRequest(requestData quotes.RequestData) ([]tinkoff.Candle, error) {
	for i := 1; i < 10; i++ {
		client := tinkoff.NewRestClient(*token)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		candles, err := client.Candles(ctx, requestData.From, requestData.To, requestData.Resolution, requestData.FIGI)

		if err == nil {
			return candles, err
		} else {
			//log.Printf("%+v\n", err)
			log.Printf("%+v\n", i)
			time.Sleep(30 * time.Second)
		}
	}
	return nil, nil
}
