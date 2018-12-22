package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Error(w http.ResponseWriter, r *http.Request) {

}

func ShowError(w http.ResponseWriter, number string) {
	errortpl, err := ioutil.ReadFile("public/templates/errors/" + number + ".html")
	if err != nil {
		fmt.Print(err)
	}

	switch number {
	case "404":
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(errortpl))
	case "500":
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(errortpl))
	case "405":
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(errortpl))
	}
}
