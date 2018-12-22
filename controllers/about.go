package controllers

import (
	"net/http"

	"../config"
)

func AboutProcess(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "about.html", nil)
}
