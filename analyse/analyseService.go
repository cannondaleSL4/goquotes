package analyse

import (
	"bytes"
	"encoding/json"
	"fmt"
	. "github.com/goquotes/constants"
	"io/ioutil"
	"net/http"
	"strings"
)

type AnalyzeRequest struct {
	Period     string `json:"period"`
	Indicator  string `json:"indicator"`
	Instrument string `json:"instrument"`
}

func GetAnalyse() {
	url := "http://localhost:5000/hello"

	analyseRequest := &AnalyzeRequest{
		Period:     "day",
		Indicator:  "rsi",
		Instrument: strings.ToLower(strings.Replace(DOWJONES, " ", "", -1)),
	}

	var jsonStr, _ = json.Marshal(analyseRequest)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
