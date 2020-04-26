package controller

import (
	tinkoff "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"github.com/goquotes/analyse"
	. "github.com/goquotes/constants"
	requestService "github.com/goquotes/request"
	"html/template"
	"log"
	"net/http"
	"time"
)

type FormAction struct {
	LastWeek, LastMonth, LastYear, For10Years, Analyse4H, AnalyseD, AnalyseW string
}

var FormActionVar = FormAction{
	LastWeek:   "Last W",
	LastMonth:  "Last M",
	LastYear:   "Last Y",
	For10Years: "For 10Y",
	Analyse4H:  "Analyse 4H",
	AnalyseD:   "Analyse D",
	AnalyseW:   "Analyse W",
}

type ViewData struct {
	Instrument    string
	Instruments   []string
	FormActionV   FormAction
	ResultUpdate  string
	ResultAnalyse []analyse.AnalyzeResponse
}

type FormData struct {
	Instruments string
	Do          string
}

func InstReloadController(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	data := ViewData{}
	if r.RequestURI == "/dj_analyse" {
		Log.Debugf("Controller for update %s instruments was chosen.", DOWJONES)
		data.Instrument = DOWJONES
		data.Instruments = GetInstrNamesDJ()
		data.FormActionV = FormActionVar
		page(w, r, data)
	} else if r.RequestURI == "/rus_analyse" {
		Log.Debugf("Controller for update %s instruments was chosen.", RUS)
		data.Instrument = RUS
		data.Instruments = GetInstrNamesRUS()
		data.FormActionV = FormActionVar
		page(w, r, data)
	} else {
		log.Fatalf("unknown instruments was chosen. %s", RUS)
	}
}

func page(w http.ResponseWriter, r *http.Request, data ViewData) {
	tmpl := template.Must(template.ParseFiles(INSTRUMENTS))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, &data)
		return
	}
	data = parseForm(r, data.Instrument, data)
	tmpl.Execute(w, &data)
}

func parseForm(r *http.Request, instr string, data ViewData) ViewData {
	var arrayOfCandle *[][]tinkoff.Candle
	fromTime := time.Now()
	if r.FormValue(FormActionVar.LastWeek) != "" {
		arrayOfCandle, _ = GetCandle(fromTime.AddDate(0, 0, -7), instr, tinkoff.CandleInterval1Hour)
		//saveToDataBase(*arrayOfCandle)
		data.ResultUpdate = "Updated"
	} else if r.FormValue(FormActionVar.LastMonth) != "" {
		arrayOfCandle, _ = GetCandle(fromTime.AddDate(0, -1, 0), instr, tinkoff.CandleInterval1Hour)
		//saveToDataBase(*arrayOfCandle)
		data.ResultUpdate = "Updated"
	} else if r.FormValue(FormActionVar.LastYear) != "" {
		arrayOfCandle, _ = GetCandle(fromTime.AddDate(0, 0, -364), instr, tinkoff.CandleInterval1Hour)
		//saveToDataBase(*arrayOfCandle)
		data.ResultUpdate = "Updated"
	} else if r.FormValue(FormActionVar.For10Years) != "" {
		arrayOfCandle, _ = GetCandle(fromTime.AddDate(0, 0, -10*364), instr, tinkoff.CandleInterval1Hour)
		//saveToDataBase(*arrayOfCandle)
		data.ResultUpdate = "Updated"
	} else if r.FormValue(FormActionVar.Analyse4H) != "" {
		var result *[]analyse.AnalyzeResponse
		arrayOfCandle, _ = GetCandle(fromTime.AddDate(0, 0, -7), instr, tinkoff.CandleInterval1Hour)
		result = analyse.GetAnalyse(arrayOfCandle, tinkoff.CandleInterval4Hour)
		if result != nil {
			data.ResultAnalyse = *result
			log.Printf("Analyse for Days has been executed.")
		}
	} else if r.FormValue(FormActionVar.AnalyseD) != "" {
		var result *[]analyse.AnalyzeResponse
		arrayOfCandle, _ = GetCandle(fromTime.AddDate(0, 0, -364), instr, tinkoff.CandleInterval1Day)
		arrayOfCandle = makeSlice(arrayOfCandle)
		result = analyse.GetAnalyse(arrayOfCandle, tinkoff.CandleInterval1Day)
		if result != nil {
			data.ResultAnalyse = *result
			log.Printf("Analyse for Days has been executed.")
		}
	} else if r.FormValue(FormActionVar.AnalyseW) != "" {
		var result *[]analyse.AnalyzeResponse
		arrayOfCandle, _ = GetCandle(fromTime.AddDate(0, -23, -20), instr, tinkoff.CandleInterval1Week)
		arrayOfCandle = makeSlice(arrayOfCandle)
		result = analyse.GetAnalyse(arrayOfCandle, tinkoff.CandleInterval1Week)
		if result != nil {
			data.ResultAnalyse = *result
			log.Printf("Analyse for Weeks has been executed.")
		}
	}
	return data
}

func makeSlice(array *[][]tinkoff.Candle) *[][]tinkoff.Candle {
	var sliceArray [][]tinkoff.Candle
	for _, element := range *array {
		temp := element
		if len(temp) > 50 {
			slice := temp[len(temp)-50:]
			sliceArray = append(sliceArray, slice)
		} else {
			log.Printf("Len of array smaller that 50 for instument %s", GetQuoteNameByFigi(temp[0].FIGI))
		}
	}
	return &sliceArray
}

func GetCandle(fromTime time.Time, instr string, interval tinkoff.CandleInterval) (*[][]tinkoff.Candle, error) {
	var err error
	var array [][]tinkoff.Candle
	if interval == tinkoff.CandleInterval1Day {
		timeArray := splitDaysData(fromTime, instr, interval)
		return getQuotes(*timeArray, fromTime, instr, interval), nil
	} else if interval == tinkoff.CandleInterval1Week {
		timeArray := splitDaysData(fromTime, instr, interval)
		return getQuotes(*timeArray, fromTime, instr, interval), nil
	} else if interval == tinkoff.CandleInterval1Hour {
		timeArray := splitHoursDate(fromTime, instr, interval)
		return getQuotes(*timeArray, fromTime, instr, interval), nil
	}

	log.Printf("unknown time period was chosen. %s", interval)
	return &array, err
}

func splitHoursDate(fromTime time.Time, instr string, interval tinkoff.CandleInterval) *[]time.Time {
	toTime := time.Now()
	var arrayOfDate []time.Time

	arrayOfDate = append(arrayOfDate, fromTime)
	arrayOfDate = append(arrayOfDate, toTime)
	return &arrayOfDate

	//var arrayOfDate []time.Time
	//toTime := time.Now()
	//
	//fromTime = roundOfTheHour(fromTime)
	//toTime = roundOfTheHour(toTime)
	//
	//diff := int(toTime.Sub(fromTime).Hours() / 24)
	//
	//for i := 0; i < diff; i = i + 6 {
	//	arrayOfDate = append(arrayOfDate, fromTime.AddDate(0, 0, +i))
	//}
	//
	//if arrayOfDate[len(arrayOfDate)-1] != toTime {
	//	arrayOfDate = append(arrayOfDate, toTime)
	//}
	//
	//return &arrayOfDate
}

func splitDaysData(fromTime time.Time, instr string, interval tinkoff.CandleInterval) *[]time.Time {
	toTime := time.Now()
	var arrayOfDate []time.Time

	arrayOfDate = append(arrayOfDate, fromTime)
	arrayOfDate = append(arrayOfDate, toTime)
	return &arrayOfDate
}

func getQuotes(arrayOfDate []time.Time, toTime time.Time, instr string, interval tinkoff.CandleInterval) *[][]tinkoff.Candle {
	var arrayOfCandle [][]tinkoff.Candle

	var quotesFIGI []string

	switch instr {
	case DOWJONES:
		quotesFIGI = GetQuotesDJ()
	case RUS:
		quotesFIGI = GetQuotesRus()
	}

	for _, element := range quotesFIGI {
		entityFromServer := requestService.UpdateFromTo(arrayOfDate, element, interval)
		if entityFromServer != nil {
			arrayOfCandle = append(arrayOfCandle, *entityFromServer)
		}
	}

	return &arrayOfCandle
}

//func saveToDataBase(entityFromServer [][]tinkoff.Candle) {
//	connection := clientDb.GetClient()
//	clientDb.InsertConst(connection)
//
//	if entityFromServer != nil {
//		clientDb.InsertNewQuotes(connection, entityFromServer)
//	}
//}

func roundOfTheHour(t time.Time) time.Time {
	year, month, day := t.Date()
	hour, _, _ := t.Clock()
	return time.Date(year, month, day, hour, 0, 0, 0, t.Location())
}
