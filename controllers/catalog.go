package controllers

import (
	"net/http"

	"../config"
	"../helpers"
	"../models"
)

type PageData struct {
	Title       string
	Description string
	Categories  []models.Cat_tree
}

//Catalog Вывод категорний и подкатегорий
func Catalog(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		helpers.ShowError(w, "405")
		return
	}

	categories, err := models.GetCategories()
	if err != nil {
		helpers.ShowError(w, "500")
		return
	}

	pagedata := PageData{Title: "Каталог - главная страница", Description: "Полезная инфо", Categories: categories}

	config.TPL.ExecuteTemplate(w, "catalog.html", pagedata)

}
