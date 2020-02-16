package controller

import (
	. "github.com/goquotes/constants"
	clientDb "github.com/goquotes/mongodbService"
	requestService "github.com/goquotes/request"
	"net/http"
	"time"
)

func RunHandler() {
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/start", indexHandlerFunc)
	http.HandleFunc("/reload", updateAllData)
	http.ListenAndServe(PORT, nil)
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index"))
}

func indexHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("start"))
}

func updateAllData(w http.ResponseWriter, r *http.Request) {
	var years = 1
	toTime := time.Now()
	fromTime := time.Now().AddDate(-years, 0, 0)
	diff := int(toTime.Sub(fromTime).Hours() / 24)
	var arrayOfDate []time.Time

	for i := 0; i < diff; i++ {
		arrayOfDate = append(arrayOfDate, toTime.AddDate(0, 0, -i))
	}

	updateDataFromTo(arrayOfDate)
}

func updateDataFromTo(dateArray []time.Time) {
	connection := clientDb.GetClient()

	for i := 1; i < len(dateArray); i++ {
		entityFromServer := requestService.UpdateFromTo(dateArray[i], dateArray[i-1])
		if entityFromServer != nil {
			clientDb.InsertNewQuotes(connection, entityFromServer)
		}
	}
}
