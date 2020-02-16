package controller

import (
	"fmt"
	. "github.com/goquotes/constants"
	"html/template"
	"log"
	"net/http"
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
	var inst string
	data := ViewData{}
	if r.RequestURI == "/dj_analyse" {
		inst = "DJ"
		Log.Debugf("Controller for update %s instruments was chosen.", inst)

		data.Instrument = "Dow Jones"
		data.Instruments = GetInstrNamesDJ()
		data.FormActionV = FormActionVar

		tmpl := template.Must(template.ParseFiles(DJ))

		if r.Method != http.MethodPost {
			fmt.Println(data)
			tmpl.Execute(w, &data)
			return
		}

		formValue := r.PostForm["form"]
		switch formValue[0] {
		case FormActionVar.LastWeek:
			fmt.Println("dsds")
		case FormActionVar.LastYear:
			fmt.Println("dsds")
		case FormActionVar.For10Years:
			fmt.Println("dsds")
		case FormActionVar.Clear:
			fmt.Println("dsds")
		case FormActionVar.Analyse:
			fmt.Println("dsds")
		}

		details := FormData{
			Instruments: r.FormValue("email"),
			Do:          r.FormValue("subject"),
		}

		// do something with details
		_ = details

		tmpl.Execute(w, &data)

		//http.HandleFunc(r.RequestURI, func(w http.ResponseWriter, r *http.Request) {
		//	if r.Method != http.MethodPost {
		//		tmpl.Execute(w, data)
		//		return
		//	}
		//
		//	details := FormData{
		//		Instruments:   r.FormValue("email"),
		//		Do: r.FormValue("subject"),
		//	}
		//
		//	// do something with details
		//	_ = details
		//
		//	tmpl.Execute(w, struct{ Success bool }{true})
		//})
	} else if r.RequestURI == "/rus_analyse" {
		inst = "RUS"
		Log.Debugf("Controller for update %s instruments was chosen.", inst)
		t, _ := template.ParseFiles(DJ)
		data.Instrument = "Rus stocks"
		t.Execute(w, data)
	} else {
		log.Fatalf("unknown instruments was chosen. %s", inst)
		t, _ := template.ParseFiles(INDEX)
		t.Execute(w, nil)
	}
}
