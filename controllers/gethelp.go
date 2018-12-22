package controllers

import (
	"net/http"

	"../config"
)

func HelpProcess(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "help.html", nil)
}
