package controllers

import (
	"net/http"

	"../config"
	"../helpers"
	"../models"
)

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

	config.TPL.ExecuteTemplate(w, "new.html", newdata)
}
