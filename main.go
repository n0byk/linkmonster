package main

import (
	"fmt"
	"log"
	"net/http"

	"./controllers"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/catalog", controllers.Catalog)
	http.HandleFunc("/show", controllers.Show)
	http.HandleFunc("/help", controllers.HelpProcess)
	http.HandleFunc("/about", controllers.AboutProcess)
	http.HandleFunc("/new", controllers.JustAdded)
	http.HandleFunc("/search", controllers.SearchProcess)

	//	get CSS & JS & ICONs
	http.HandleFunc("/public/static/css/skeleton.css", controllers.GetCssSkeleton)
	http.HandleFunc("/public/static/css/normalize.css", controllers.GetCssNormalize)
	http.HandleFunc("/favicon.ico", controllers.GetFavIcon)

	//	ERORS
	http.HandleFunc("/error", controllers.Error)

	//	START SERVER
	//go http.ListenAndServe(":8080", nil)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Println("server started  :8080.")
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/catalog", http.StatusSeeOther)
}
