package controller

import (
	"fmt"
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
	LastWeek, LastYear, For10Years, Clear, Analyse string
}

var FormActionVar = FormAction{
	LastWeek:   "Last week",
	LastYear:   "Last year",
	For10Years: "For 10 years",
	Clear:      "Clear",
	Analyse:    "Analyse",
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
	if r.FormValue(FormActionVar.LastWeek) != "" {
		fromTime := time.Now().AddDate(0, 0, -7)
		updateData(fromTime, instr)
	} else if r.FormValue(FormActionVar.LastYear) != "" {
		fromTime := time.Now().AddDate(-1, 0, 0)
		updateData(fromTime, instr)
	} else if r.FormValue(FormActionVar.For10Years) != "" {
		fromTime := time.Now().AddDate(-10, 0, 0)
		updateData(fromTime, instr)
	} else if r.FormValue(FormActionVar.Clear) != "" {
		fmt.Println("dsds")
	} else if r.FormValue(FormActionVar.Analyse) != "" {
		analyse.GetAnalyse()
	}
}

func updateData(fromTime time.Time, instr string) {
	toTime := time.Now()
	diff := int(toTime.Sub(fromTime).Hours() / 24)
	var arrayOfDate []time.Time

	for i := 0; i < diff; i++ {
		arrayOfDate = append(arrayOfDate, toTime.AddDate(0, 0, -i))
	}
	updateDataFromTo(arrayOfDate, instr)
	Log.Debugf("Request for update %s from %s to %s")
}

func updateDataFromTo(dateArray []time.Time, instr string) {
	connection := clientDb.GetClient()
	clientDb.InsertConst(connection)
	for i := 1; i < len(dateArray); i++ {
		entityFromServer := requestService.UpdateFromTo(dateArray[i], dateArray[i-1], instr)
		if entityFromServer != nil {
			clientDb.InsertNewQuotes(connection, entityFromServer)
		}
	}
}
