package controller

import (
	tinkoff "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"github.com/goquotes/analyse"
	. "github.com/goquotes/constants"
	clientDb "github.com/goquotes/mongodbService"
	requestService "github.com/goquotes/request"
	"html/template"
	"log"
	"net/http"
	"time"
)

type FormAction struct {
	LastWeek, LastMonth, LastYear, For10Years, Analyse4, AnalyseD string
}

var FormActionVar = FormAction{
	LastWeek:   "Last W",
	LastMonth:  "Last M",
	LastYear:   "Last Y",
	For10Years: "For 10Y",
	Analyse4:   "Analyse 4H",
	AnalyseD:   "Analyse D",
}

type ViewData struct {
	Instrument  string
	Instruments []string
	FormActionV FormAction
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

		tmpl := template.Must(template.ParseFiles(INSTRUMENTS))

		if r.Method != http.MethodPost {
			tmpl.Execute(w, &data)
			return
		}

		parseForm(r, DOWJONES)

		details := FormData{
			Instruments: r.FormValue("email"),
			Do:          r.FormValue("subject"),
		}
		_ = details

		tmpl.Execute(w, &data)

	} else if r.RequestURI == "/rus_analyse" {
		Log.Debugf("Controller for update %s instruments was chosen.", RUS)
		data.Instrument = RUS
		data.Instruments = GetInstrNamesRUS()
		data.FormActionV = FormActionVar
		t, _ := template.ParseFiles(INSTRUMENTS)
		data.Instrument = "Rus stocks"
		t.Execute(w, data)
	} else {
		log.Fatalf("unknown instruments was chosen. %s", RUS)
	}
}

func parseForm(r *http.Request, instr string) {
	var arrayOfCandle []tinkoff.Candle
	fromTime := time.Now()
	if r.FormValue(FormActionVar.LastWeek) != "" {
		arrayOfCandle = getCandle(fromTime.AddDate(0, 0, -7), instr, tinkoff.CandleInterval1Hour)
		saveToDataBase(arrayOfCandle)
	} else if r.FormValue(FormActionVar.LastMonth) != "" {
		arrayOfCandle = getCandle(fromTime.AddDate(0, -1, 0), instr, tinkoff.CandleInterval1Hour)
		saveToDataBase(arrayOfCandle)
	} else if r.FormValue(FormActionVar.LastYear) != "" {
		arrayOfCandle = getCandle(fromTime.AddDate(0, 0, -364), instr, tinkoff.CandleInterval1Hour)
		saveToDataBase(arrayOfCandle)
	} else if r.FormValue(FormActionVar.For10Years) != "" {
		arrayOfCandle = getCandle(fromTime.AddDate(0, 0, -10*364), instr, tinkoff.CandleInterval1Hour)
		saveToDataBase(arrayOfCandle)
	} else if r.FormValue(FormActionVar.Analyse4) != "" {
		arrayOfCandle = getCandle(fromTime.AddDate(0, 0, -7), instr, tinkoff.CandleInterval1Hour)
		analyse.GetAnalyse(arrayOfCandle, 4)
	} else if r.FormValue(FormActionVar.AnalyseD) != "" {
		arrayOfCandle = getCandle(fromTime.AddDate(0, 0, -364), instr, tinkoff.CandleInterval1Day)
		analyse.GetAnalyse(arrayOfCandle, 24)
	}
}

func getCandle(fromTime time.Time, instr string, interval tinkoff.CandleInterval) []tinkoff.Candle {

	if interval == tinkoff.CandleInterval1Day {
		return splitDaysData(fromTime, instr, interval)
	} else {
		return splitHoursDate(fromTime, instr, interval)
	}
}

func splitHoursDate(fromTime time.Time, instr string, interval tinkoff.CandleInterval) []tinkoff.Candle {
	var arrayOfCandle []tinkoff.Candle
	var arrayOfDate []time.Time
	toTime := time.Now()

	fromTime = roundOfTheHour(fromTime)
	toTime = roundOfTheHour(toTime)

	diff := int(toTime.Sub(fromTime).Hours() / 24)

	//arrayOfDate = append(arrayOfDate, toTime)

	for i := 0; i < diff; i = i + 6 {
		arrayOfDate = append(arrayOfDate, fromTime.AddDate(0, 0, +i))
	}

	if arrayOfDate[len(arrayOfDate)-1] != toTime {
		arrayOfDate = append(arrayOfDate, toTime)
	}

	for i := 1; i < len(arrayOfDate); i++ {
		if checkWeekEnd(arrayOfDate[i], arrayOfDate[i-1]) {
			entityFromServer := requestService.UpdateFromTo(arrayOfDate[i-1], arrayOfDate[i], instr, interval)
			if entityFromServer != nil {
				arrayOfCandle = append(arrayOfCandle, entityFromServer...)
			}
		}
	}
	return arrayOfCandle
}

func splitDaysData(fromTime time.Time, instr string, interval tinkoff.CandleInterval) []tinkoff.Candle {
	toTime := time.Now()
	var arrayOfDate []time.Time

	fromTime = roundOfTheDay(fromTime)
	toTime = roundOfTheDay(toTime)

	arrayOfDate = append(arrayOfDate, fromTime)
	arrayOfDate = append(arrayOfDate, toTime)

	entityFromServer := requestService.UpdateFromTo(fromTime, toTime, instr, interval)
	if entityFromServer != nil {
		return entityFromServer
	}
	return nil
}

func saveToDataBase(entityFromServer []tinkoff.Candle) {
	connection := clientDb.GetClient()
	clientDb.InsertConst(connection)

	if entityFromServer != nil {
		clientDb.InsertNewQuotes(connection, entityFromServer)
	}
}

func roundOfTheHour(t time.Time) time.Time {
	year, month, day := t.Date()
	hour, _, _ := t.Clock()
	return time.Date(year, month, day, hour, 0, 0, 0, t.Location())
}

func roundOfTheDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func checkWeekEnd(from time.Time, to time.Time) bool {
	if (from.Weekday() != 6 && to.Weekday() != 0) && (from.Weekday() != 0 && to.Weekday() != 1) {
		return true
	}
	return false
}
