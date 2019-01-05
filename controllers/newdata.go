package controllers

import (
	"net/http"

	"../config"
	"../helpers"
	"../models"
)

type PageDataNew struct {
	Title       string
	Description string
	Newdata     []models.GetifobyUID
}

// JustAdded Получение последних N  добавленных записей DEFAULT == 20
func JustAdded(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		helpers.ShowError(w, "405")
		return
	}

	newdata, err := models.GetNewData(20)
	if err != nil {
		helpers.ShowError(w, "404")
		return
	}
	pagedatanew := PageDataNew{Title: "Каталог - главная страница", Description: "Полезная инфо", Newdata: newdata}
	config.TPL.ExecuteTemplate(w, "new.html", pagedatanew)
}
