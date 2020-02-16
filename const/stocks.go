package _const

import (
	"strconv"
	"time"
)

type stringTime string

// this struct is data from response
type StocksFromResponse struct {
	Code  string     `json:"code"`
	Open  float64    `json:"o"`
	High  float64    `json:"h"`
	Low   float64    `json:"l"`
	Close float64    `json:"c"`
	Vol   float64      `json:"v"`
	Date  stringTime `json:"date"`
}

// overrided func for unamarshal
func (j *stringTime) UnmarshalJSON(data []byte) error {
	millis, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	*j = stringTime(time.Unix(millis, 0).String())
	return nil
}

