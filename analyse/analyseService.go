package analyse

import (
	"fmt"
	tinkoff "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"github.com/goquotes/constants"
	"github.com/markcheno/go-talib"
	"github.com/sdcoffey/big"
	"github.com/sdcoffey/techan"
	"log"
	"time"
)

type AnalyzeResponse struct {
	Interval    string `json:interval`
	Name        string `json:"name"`
	Indicator   string `json:"indicator"`
	Result      string `json:"result"`
	Description string `json:"description"`
}

func GetAnalyse(arrayOfQuotes *[][]tinkoff.Candle, interval tinkoff.CandleInterval) *[]AnalyzeResponse {
	var hours int
	if interval == tinkoff.CandleInterval1Day {
		hours = 24
	} else if interval == tinkoff.CandleInterval4Hour {
		hours = 4
	} else if interval == tinkoff.CandleInterval1Week {
		hours = 24 * 7
	}

	var results []AnalyzeResponse
	for _, element := range *arrayOfQuotes {
		series := techan.NewTimeSeries()
		for _, innerElement := range element {
			period := techan.NewTimePeriod(time.Unix(innerElement.TS.Unix(), 0), time.Hour*time.Duration(hours))
			candle := techan.NewCandle(period)
			candle.OpenPrice = big.NewDecimal(innerElement.OpenPrice)
			candle.ClosePrice = big.NewDecimal(innerElement.ClosePrice)
			candle.MaxPrice = big.NewDecimal(innerElement.HighPrice)
			candle.MinPrice = big.NewDecimal(innerElement.LowPrice)
			series.AddCandle(candle)
		}

		var result *AnalyzeResponse
		var resultWarning *AnalyzeResponse
		result = getRsi(*series, element[0].FIGI, interval)
		resultWarning = getRsiWarning(*series, element[0].FIGI, interval)
		//getWilliams(*series, element[0].FIGI, interval)
		if result != nil {
			results = append(results, *result)
		}

		if resultWarning != nil {
			results = append(results, *resultWarning)
		}
	}

	return &results
}

func getRsi(series techan.TimeSeries, name string, interval tinkoff.CandleInterval) *AnalyzeResponse {
	name = constants.GetQuoteNameByFigi(name)
	var arrayClose []float64

	for _, element := range series.Candles {
		arrayClose = append(arrayClose, element.ClosePrice.Float())
	}
	rsi := talib.Rsi(arrayClose, 14)
	slice_rsi := rsi[len(rsi)-5:]

	was_in_down := false
	for _, price := range slice_rsi {
		if price < 30 {
			was_in_down = true
		}

		if was_in_down {
			preLast := slice_rsi[len(slice_rsi)-2]
			last := slice_rsi[len(slice_rsi)-1]

			if preLast < 30 && last > preLast && last > 30 && getWilliams(series) {
				var result AnalyzeResponse
				result.Indicator = "Rsi"
				result.Interval = string(interval)
				result.Name = fmt.Sprintf("(%s) %s", constants.GetFigiByName(name), cutName(name))
				result.Result = "Buy"
				result.Description = fmt.Sprintf("preRsi: %d , lastRsi: %d", int64(preLast), int(last))
				log.Printf("result of analyse for indicator %s, for instrument %s . preRsi: %f , lastRsi: %f , result: %s", "Rsi",
					name, preLast, last, "Buy")
				return &result
			}
		}
	}
	return nil
}

func getRsiWarning(series techan.TimeSeries, name string, interval tinkoff.CandleInterval) *AnalyzeResponse {
	name = constants.GetQuoteNameByFigi(name)
	var arrayClose []float64

	for _, element := range series.Candles {
		arrayClose = append(arrayClose, element.ClosePrice.Float())
	}
	rsi := talib.Rsi(arrayClose, 14)
	slice_rsi := rsi[len(rsi)-5:]
	last := slice_rsi[len(slice_rsi)-1]
	if last < 30 {
		var result AnalyzeResponse
		result.Indicator = "Rsi"
		result.Interval = string(interval)
		result.Name = fmt.Sprintf("(%s) %s", constants.GetFigiByName(name), cutName(name))
		result.Result = "Buy Warning"
		result.Description = fmt.Sprintf("last RSI :%d", int64(last))
		return &result
	}
	return nil
}

func cutName(name string) string {
	if len(name) > 15 {
		return name[:15]
	}
	return name
}

func getWilliams(series techan.TimeSeries) bool {
	var arrayClose []float64
	var arrayLow []float64
	var arrayHigh []float64
	for _, element := range series.Candles {
		arrayClose = append(arrayClose, element.ClosePrice.Float())
		arrayLow = append(arrayClose, element.MinPrice.Float())
		arrayHigh = append(arrayClose, element.MaxPrice.Float())
	}
	williams := talib.WillR(arrayHigh, arrayLow, arrayClose, 14)
	if williams[len(williams)-1] > -80 && williams[len(williams)-1] < -20 {
		return true
	}
	return false
}
