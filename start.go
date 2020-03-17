package main

import (
	"flag"
	. "github.com/goquotes/constants"
	controller "github.com/goquotes/controller"
	. "github.com/goquotes/scheduler"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {

	go func() {
		Scheduler()
	}()
	port := flag.String("port", os.Getenv("PORT"), "app port")
	if len(*port) == 0 {
		*port = "3000"
	}
	Log.Infof("Started DNS MX Record Application. URL Port [%v] ", *port)
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", controller.StartHandler)
	rtr.HandleFunc("/dj_analyse", controller.InstReloadController)
	rtr.HandleFunc("/rus_analyse", controller.InstReloadController)
	rtr.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.Handle("/", rtr)
	Log.Fatal(http.ListenAndServe(":"+*port, nil))
}
