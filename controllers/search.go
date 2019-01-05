package controllers

import (
	"net/http"

	"../config"
	"../helpers"
)

//  Поиск
func SearchProcess(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		helpers.ShowError(w, "405")
		return
	}

	pagedatasingle := PageDataSingle{Title: "Каталог - главная страница", Description: "Полезная инфо"}
	config.TPL.ExecuteTemplate(w, "search.html", pagedatasingle)

}
