package main

import (
	. "github.com/goquotes/constants"
	controller "github.com/goquotes/controller"
	. "github.com/goquotes/scheduler"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	go func() {
		Scheduler()
	}()
	Log.Infof("Started DNS MX Record Application. URL Port [%v] ", PORT)
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", controller.StartHandler)
	rtr.HandleFunc("/dj_analyse", controller.InstReloadController)
	rtr.HandleFunc("/rus_analyse", controller.InstReloadController)
	rtr.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.Handle("/", rtr)
	Log.Fatal(http.ListenAndServe(PORT, nil))

}
