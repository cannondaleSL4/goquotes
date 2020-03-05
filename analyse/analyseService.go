package analyse

import (
	"fmt"
	tinkoff "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"github.com/sdcoffey/big"
	"strconv"
	"time"

	"github.com/sdcoffey/techan"
)

type AnalyzeRequest struct {
	Interval   string `json:interval`
	Period     string `json:"period"`
	Indicator  string `json:"indicator"`
	Instrument string `json:"instrument"`
}

func GetAnalyse(arrayOfQuotes []tinkoff.Candle, hours int) {

	series := techan.NewTimeSeries()

	//dataSet := [][]string{}
	//
	//for _,element := range arrayOfQuotes {
	//	var innerElement []string
	//	innerElement = append(innerElement,string(element.TS.Unix()))
	//	innerElement = append(innerElement,string(element.OpenPrice))
	//}

	//dataset := [][]string{
	//	// Timestamp, Open, Close, High, Low, volume
	//	{"1582232400", "320.300000", "319.780000", "320.500000", "319.770000", "1259631"},
	//	{"1582236000", "319.800000", "319.460000", "320.070000", "319.460000", "4162"},
	//	{"1582268400", "318.040000", "317.820000", "320.230000", "315.000000", "2623"},
	//	{"1582272000", "317.800000", "317.520000", "317.980000", "315.850000", "1201"},
	//	{"1582275600", "317.730000", "317.330000", "318.340000", "317.310000", "439"},
	//	{"1582279200", "317.690000", "318.360000", "318.360000", "317.550000", "244"},
	//	{"1582282800", "318.170000", "318.300000", "318.600000", "317.800000", "196"},
	//	{"1582286400", "318.200000", "318.750000", "319.160000", "318.030000", "3319"},
	//	{"1582290000", "318.900000", "318.370000", "319.000000", "318.200000", "15922"},
	//	{"1582293600", "318.370000", "317.270000", "320.450000", "316.300000", "2210154"},
	//	{"1582297200", "317.230000", "316.290000", "317.770000", "315.600000", "1192793"},
	//	{"1582300800", "316.300000", "316.930000", "317.280000", "315.650000", "614976"},
	//	{"1582304400", "316.970000", "315.100000", "317.420000", "314.680000", "621007"},
	//	{"1582308000", "315.100000", "312.110000", "315.490000", "311.180000", "1116452"},
	//	{"1582311600", "312.080000", "312.210000", "312.630000", "310.500000", "976790"},
	//	{"1582315200", "312.130000", "313.030000", "313.530000", "310.840000", "1363454"},
	//	{"1582527600", "307.000000", "290.090000", "309.070000", "290.000000", "5185"},
	//	{"1582531200", "292.000000", "305.000000", "305.100000", "290.090000", "2417"},
	//	{"1582534800", "304.800000", "299.680000", "305.000000", "297.910000", "3328"},
	//	{"1582538400", "299.680000", "298.870000", "301.090000", "298.520000", "2454"},
	//	{"1582542000", "299.530000", "300.260000", "301.200000", "298.370000", "1598"},
	//	{"1582545600", "300.310000", "301.930000", "303.880000", "300.310000", "30222"},
	//	{"1582549200", "301.930000", "301.250000", "302.220000", "300.210000", "139740"},
	//	{"1582552800", "301.370000", "302.760000", "304.120000", "289.230000", "5997405"},
	//	{"1582556400", "302.770000", "301.120000", "303.690000", "298.400000", "2028752"},
	//	{"1582560000", "301.080000", "299.730000", "301.340000", "298.410000", "1138451"},
	//	{"1582563600", "299.660000", "300.120000", "300.970000", "298.000000", "975820"},
	//	{"1582567200", "300.070000", "300.370000", "300.670000", "298.660000", "775653"},
	//	{"1582570800", "300.340000", "302.180000", "302.500000", "299.100000", "960772"},
	//	{"1582574400", "302.200000", "298.170000", "303.280000", "298.090000", "2304478"},
	//}

	for _, element := range arrayOfQuotes {
		//start, _ := strconv.ParseInt(element.TS.Unix(), 10, 64)
		period := techan.NewTimePeriod(time.Unix(element.TS.Unix(), 0), time.Hour*time.Duration(hours))
		candle := techan.NewCandle(period)
		candle.OpenPrice = big.NewDecimal(element.OpenPrice)
		candle.ClosePrice = big.NewDecimal(element.ClosePrice)
		candle.MaxPrice = big.NewDecimal(element.HighPrice)
		candle.MinPrice = big.NewDecimal(element.LowPrice)

		series.AddCandle(candle)
	}

	for _, datum := range dataset {
		start, _ := strconv.ParseInt(datum[0], 10, 64)
		period := techan.NewTimePeriod(time.Unix(start, 0), time.Hour*time.Duration(hours))

		candle := techan.NewCandle(period)
		candle.OpenPrice = big.NewFromString(datum[1])
		candle.ClosePrice = big.NewFromString(datum[2])
		candle.MaxPrice = big.NewFromString(datum[3])
		candle.MinPrice = big.NewFromString(datum[4])

		series.AddCandle(candle)
	}

	closePrices := techan.NewClosePriceIndicator(series)
	movingAverage := techan.NewEMAIndicator(closePrices, 10) // Create an exponential moving average with a window of 10

	fmt.Println(movingAverage.Calculate(0).FormattedString(2))
}
