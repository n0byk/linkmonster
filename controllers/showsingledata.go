package controllers

import (
	"database/sql"
	"net/http"

	"../config"
	"../models"
)

//  Show Получение информации с конкретным ID
func Show(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	singledata, err := models.ShowSingleData(r)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		ShowError(w, "404")

		//http.Redirect(w, r, "/404", http.StatusTemporaryRedirect)
		return
	}
	config.TPL.ExecuteTemplate(w, "showsingledata.html", singledata)
}
