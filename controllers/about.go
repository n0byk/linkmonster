package controllers

//Контроллер отдает страницу о провекте
import (
	"net/http"

	"../config"
)

func AboutProcess(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "about.html", nil)
}
