package controllers

import (
	"fmt"
	"net/http"

	"../config"
	"../helpers"
)

// GetAllFromCat Получение последних N  добавленных записей DEFAULT == 20
func GetAllFromCat(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		helpers.ShowError(w, "405")
		return
	}

	p := helpers.NewPagination(453, 33, 5, 5)

	fmt.Printf("%+v\n", p)

	config.TPL.ExecuteTemplate(w, "getallfromcat.html", p)
}
