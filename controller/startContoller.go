package controller

import (
	. "github.com/goquotes/constants"
	"html/template"
	"net/http"
)

func StartHandler(w http.ResponseWriter, r *http.Request) {

	Log.Debug("RootHandler started and redirecting to the index page")

	t, _ := template.ParseFiles(INDEX)
	t.Execute(w, nil)
}
