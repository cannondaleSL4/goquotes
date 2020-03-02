package request

import (
	"context"
	"flag"
	. "github.com/goquotes/constants"
	"log"
	"os"
	"time"

	tinkoff "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
)

var token = flag.String("token", os.Getenv("TOKEN"), "your token")

func UpdateFromTo(from time.Time, to time.Time, instr string) []tinkoff.Candle {
	var quotesFIGI []string
	switch instr {
	case DOWJONES:
		quotesFIGI = GetQuotesDJ()
	case RUS:
		quotesFIGI = GetQuotesRus()
	}
	var arrayOfRequestData []RequestData
	for _, element := range quotesFIGI {
		var req RequestData
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

func requestToServer(arrayOfRequestData []RequestData) []tinkoff.Candle {
	stocks := make([]tinkoff.Candle, 0)

	for _, data := range arrayOfRequestData {
		candles, err := makeRequest(data)
		if err != nil {
			log.Printf("%+v\n", err)
			return nil
		}

		if len(candles) == 0 {
			Log.Debugf("Len of the response array is 0 for data: %+v\n", data.From)
			Log.Debugf("Request for instrument %+v\n", data.FIGI)
			return nil
		}
		stocks = append(stocks, candles...)
	}

	return stocks
}

func makeRequest(requestData RequestData) ([]tinkoff.Candle, error) {
	for i := 1; i < 10; i++ {
		client := tinkoff.NewRestClient(*token)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		candles, err := client.Candles(ctx, requestData.From, requestData.To, requestData.Resolution, requestData.FIGI)
		ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err == nil {
			if len(candles) != 0 {
				return candles, err
			} else {
				Log.Debugf("sleep for %d seconds", 5)
				Log.Debugf("date from %+v to %+v", requestData.From, requestData.To)
				time.Sleep(5 * time.Second)
			}
		} else {
			Log.Debugf("sleep for %d seconds", 30)
			time.Sleep(30 * time.Second)
		}
	}
	return nil, nil
}
