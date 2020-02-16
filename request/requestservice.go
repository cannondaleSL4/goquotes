package request

import (
	"bytes"
	"encoding/json"
	quotes "github.com/goquotes/const"

	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)


func UpdateFromTo(from string, to string) []quotes.StocksFromResponse{
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

func requestToServer(arrayOfRequestData []quotes.RequestData) []quotes.StocksFromResponse{
	stocks := make([]quotes.StocksFromResponse,0)

	for _, element := range arrayOfRequestData {
		convertJson, _ := json.Marshal(element)
		jsonStr := []byte(convertJson)
		req, err := http.NewRequest("POST", quotes.URL, bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body) // []byte
		value := gjson.Get(string(body), "payload.candles")

		json.Unmarshal([]byte(value.String()), &stocks)
		for x := range stocks {
			stocks[x].Code = element.Ticker
		}
		//clientDb.InsertNewQuotes(connection, stocks)
	}
	return stocks
}

func insertToDatabase(arrayOfRequestData []quotes.RequestData) {
	//connection := clientDb.GetClient()

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
	//	stocks := make([]quotes.StocksFromResponse,0)
	//	json.Unmarshal([]byte(value.String()), &stocks)
	//	for x := range stocks {
	//		stocks[x].Code = element.Ticker
	//	}
	//	clientDb.InsertNewQuotes(connection, stocks)
	//}
}

