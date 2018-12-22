package controllers

import (
	"net/http"
)

func GetCssNormalize(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/static/css/normalize.css")
}
func GetCssSkeleton(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/static/css/skeleton.css")
}
func GetFavIcon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/static/img/favicon.ico")
}
