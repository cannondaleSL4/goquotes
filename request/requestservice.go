package request

import (
	"context"
	"flag"
	"log"

	. "github.com/goquotes/constants"
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
			log.Printf("Len of the reponse array is 0 for data: %+v\n", data.From)
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
		defer cancel()
		candles, err := client.Candles(ctx, requestData.From, requestData.To, requestData.Resolution, requestData.FIGI)

		if err == nil {
			return candles, err
		} else {
			Log.Debugf("%+v\n", i)
			time.Sleep(30 * time.Second)
		}
	}
	return nil, nil
}
