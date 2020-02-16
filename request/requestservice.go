package request

import (
	"context"
	"flag"
	"fmt"
	quotes "github.com/goquotes/const"
	"log"
	"os"
	"time"

	tinkoff "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
)

var token = flag.String("token", os.Getenv("TOKEN"), "your token")

func UpdateFromTo(from string, to string) []quotes.StocksFromResponse {
	var arrayOfRequestData []quotes.RequestData
	for _, element := range quotes.GetQuotes() {
		var req quotes.RequestData
		req.Ticker = element
		req.From = from
		req.To = to
		req.Resolution = "D"
		arrayOfRequestData = append(arrayOfRequestData, req)
	}
	return requestToServer(arrayOfRequestData)
}

func requestToServer(arrayOfRequestData []quotes.RequestData) []quotes.StocksFromResponse {
	stocks := make([]quotes.StocksFromResponse, 0)

	client := tinkoff.NewRestClient(*token)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, element := range arrayOfRequestData {
		instruments, err := client.SearchInstrumentByTicker(ctx, element.Ticker)
		if err != nil {
			log.Fatalln(err)
		}

		if len(instruments) > 1 {
			log.Fatalf("server return more one instrument %+v\n", instruments)
		}

		for _, instrument := range instruments {
			fmt.Printf("{\"FIGI\":\"%+v\",\"Tiker\":\"%+v\",\"Name\":\"%+v\"},\n", instrument.FIGI, instrument.Ticker, instrument.Name)
		}

		//log.Printf("%+v\n", instruments)

		ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
	}

	//instruments, err := client.SearchInstrumentByTicker(ctx, "NKE")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//if len(instruments) > 1 {
	//	log.Fatalf("server return more one instrument %+v\n", instruments)
	//}
	//
	//log.Printf("%+v\n", instruments)
	//
	//ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()

	//for _, element := range arrayOfRequestData {
	//	convertJson, _ := json.Marshal(element)
	//	jsonStr := []byte(convertJson)
	//	req, err := http.NewRequest("POST", quotes.URL, bytes.NewBuffer(jsonStr))
	//	req.Header.Set("Content-Type", "application/json")
	//
	//	client := &http.Client{}
	//	resp, err := client.Do(req)
	//	if err != nil {
	//		panic(err)
	//	}
	//	defer resp.Body.Close()
	//
	//	body, _ := ioutil.ReadAll(resp.Body) // []byte
	//	value := gjson.Get(string(body), "payload.candles")
	//
	//	json.Unmarshal([]byte(value.String()), &stocks)
	//	for x := range stocks {
	//		stocks[x].Code = element.Ticker
	//	}
	//}
	return stocks
}
