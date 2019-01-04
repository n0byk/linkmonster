package controllers

import (
	"database/sql"
	"net/http"

	"../config"
	"../helpers"
	"../models"
)

type PageDataSingle struct {
	Title       string
	Description string
	Singledata  models.GetifobyUID
}

// Show Получение информации с конкретным ID
func Show(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		helpers.ShowError(w, "405")
		return
	}

	singledata, err := models.ShowSingleData(w, r)

	if err == sql.ErrNoRows || err != nil {
		helpers.ShowError(w, "404")
		return
	}

	pagedatasingle := PageDataSingle{Title: "Каталог - главная страница", Description: "Полезная инфо", Singledata: singledata}
	config.TPL.ExecuteTemplate(w, "showsingledata.html", pagedatasingle)
}
