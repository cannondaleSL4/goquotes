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
	Interval  string `json:interval`
	Name      string `json:"name"`
	Indicator string `json:"indicator"`
	Result    string `json:"result"`
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
		result = getRsi(*series, element[0].FIGI, interval)
		//is_best_moment := getPoint(close[len(close)-5:],element[0].FIGI, interval)
		is_best_moment := getPoint(*series, element[0].FIGI, interval)
		_ = is_best_moment
		//getWilliams(*series, element[0].FIGI, interval)
		if result != nil {
			results = append(results, *result)
		}
	}

	return &results
}

func getWilliams(series techan.TimeSeries, name string, interval tinkoff.CandleInterval) {
	var arrayClose []float64
	var arrayLow []float64
	var arrayHigh []float64
	for _, element := range series.Candles {
		arrayClose = append(arrayClose, element.ClosePrice.Float())
		arrayLow = append(arrayClose, element.MinPrice.Float())
		arrayHigh = append(arrayClose, element.MaxPrice.Float())
	}
	rsi2 := talib.WillR(arrayHigh, arrayLow, arrayClose, 14)

	_ = rsi2
}

func getRsi(series techan.TimeSeries, name string, interval tinkoff.CandleInterval) *AnalyzeResponse {
	var arrayClose []float64

	for _, element := range series.Candles {
		arrayClose = append(arrayClose, element.ClosePrice.Float())
	}
	rsi := talib.Rsi(arrayClose, 14)

	pre := rsi[len(rsi)-3]
	current := rsi[len(rsi)-2]

	preLineDown := pre > 80
	afterLineDown := current < 80
	preLineUp := pre < 20
	afterLineUp := current > 20

	//predict for feature
	featureDown := current > 80
	featureUp := current < 20

	var result AnalyzeResponse

	nameInstrument := constants.GetQuoteNameByFigi(name)

	if preLineDown && afterLineDown {
		result.Indicator = "Rsi"
		result.Interval = string(interval)
		result.Name = nameInstrument
		result.Result = "Sell"
		log.Printf("result of analyse for indicator %s, for instrument %s . preRsi: %f , lastRsi: %f , result: %f", "Rsi",
			nameInstrument, pre, current, "Sell")
		return &result
	} else if preLineUp && afterLineUp {
		result.Indicator = "Rsi"
		result.Interval = string(interval)
		result.Name = nameInstrument
		result.Result = "Buy"
		log.Printf("result of analyse for indicator %s, for instrument %s . preRsi: %f , lastRsi: %f , result: %s", "Rsi",
			nameInstrument, pre, current, "Buy")
		return &result
	} else if featureDown || featureUp {
		log.Printf("*************instrument %s has lastRsi %f ****************", nameInstrument, current)
	}

	return nil
}

func getPoint(series techan.TimeSeries, name string, interval tinkoff.CandleInterval) bool {
	name = constants.GetQuoteNameByFigi(name)
	var arrayClose []float64

	for _, element := range series.Candles {
		arrayClose = append(arrayClose, element.ClosePrice.Float())
	}
	rsi := talib.Rsi(arrayClose, 14)
	slice_rsi := rsi[len(rsi)-5:]

	was_in_down := false
	for _, price := range slice_rsi {
		if price < 20 {
			was_in_down = true
		}

		if was_in_down {
			preLast := slice_rsi[len(slice_rsi)-2]
			last := slice_rsi[len(slice_rsi)-1]

			if last > preLast && last > 30 {
				fmt.Printf("how best moment for buy %s in inteval %s, RSI is %f \n", name, interval, last)
				return true
			}
		}
	}
	return false
}
