package controller

import (
	clientDb "github.com/goquotes/mongodbService"
	requestService "github.com/goquotes/request"
	"net/http"
	"time"
)

func RunHandler() {
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/start", indexHandlerFunc)
	http.HandleFunc("/reload", reloadData)
	http.ListenAndServe(":3000", nil)
}

func indexPage(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("index"))
}

func indexHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("start"))
}

func reloadData(w http.ResponseWriter, r *http.Request) {
	connection := clientDb.GetClient()
	var years = 1
	var arrayOfTime [] string

	toTime := time.Now().Format(time.RFC3339)

	for i := years; i >= 1; i-- {
		fromTime := time.Now().AddDate(-i, 0, 0).Format(time.RFC3339)
		arrayOfTime = append(arrayOfTime, fromTime)
	}

	arrayOfTime = append(arrayOfTime, toTime)

	for i := 0; i<len(arrayOfTime)-1 ;i++ {
		entityFromServer :=	requestService.UpdateFromTo(arrayOfTime[i], arrayOfTime[i+1])
		clientDb.InsertNewQuotes(connection, entityFromServer)
	}
}
