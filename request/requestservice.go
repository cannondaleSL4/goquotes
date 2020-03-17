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

func UpdateFromTo(timesArray []time.Time, figi string, interval tinkoff.CandleInterval) *[]tinkoff.Candle {
	var candlesArray []tinkoff.Candle
	for i := 1; i < len(timesArray); i++ {
		var req RequestData
		req.FIGI = figi
		req.From = timesArray[i-1]
		req.To = timesArray[i]
		req.Resolution = interval
		resp := requestToServer(req)
		if resp != nil && len(*resp) != 0 {
			candlesArray = append(candlesArray, *resp...)
		}
	}
	if len(candlesArray) != 0 {
		return &candlesArray
	}
	return nil
}

func requestToServer(requestData RequestData) *[]tinkoff.Candle {
	stocks := make([]tinkoff.Candle, 0)

	candles, err := makeRequest(requestData)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil
	}

	if len(candles) == 0 {
		Log.Debugf("Len of the response array is 0 for data: %+v\n", requestData.From)
		Log.Debugf("Request for instrument %+v\n", requestData.FIGI)
		return nil
	}
	stocks = append(stocks, candles...)

	return &stocks
}

func makeRequest(requestData RequestData) ([]tinkoff.Candle, error) {
	for i := 1; i < 5; i++ {
		client := tinkoff.NewRestClient(*token)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		candles, err := client.Candles(ctx, requestData.From, requestData.To, requestData.Resolution, requestData.FIGI)
		ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err == nil {
			if len(candles) != 0 {
				return candles, err
			} else {
				Log.Debugf("sleep for %d milliseconds", 500)
				Log.Debugf("date from %+v to %+v", requestData.From, requestData.To)
				time.Sleep(500 * time.Millisecond)
			}
		} else {
			Log.Debugf("error: %s ", err)
			Log.Debugf("sleep for %d seconds", 30)
			time.Sleep(30 * time.Second)
		}
	}
	return nil, nil
}
