package controllers

import (
	"net/http"

	"../config"
	"../models"
)

type PageData struct {
	Title       string
	Description string
	Categories  []models.Cat_tree
}

//Вывод категорний и подкатегорий
func Catalog(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	categories, err := models.GetCategories()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	pagedata := PageData{Title: "Каталог - главная страница", Description: "Полезная инфо", Categories: categories}

	config.TPL.ExecuteTemplate(w, "catalog.html", pagedata)

}
