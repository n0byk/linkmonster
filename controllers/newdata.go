package controllers

import (
	"net/http"

	"../config"
	"../models"
)

// JustAdded Получение последних N  добавленных записей DEFAULT == 20
func JustAdded(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	newdata, err := models.GetNewData(20)
	if err != nil {
		http.Redirect(w, r, "/404", http.StatusNotFound)
		return
	}

	config.TPL.ExecuteTemplate(w, "new.html", newdata)
}
