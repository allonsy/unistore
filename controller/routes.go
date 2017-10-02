package controller

import (
	"net/http"
)

func Init() {
	http.HandleFunc("/", getHome)
	http.HandleFunc("/html/", handleSource)
	http.HandleFunc("/css/", handleSource)
	http.HandleFunc("/js/", handleJS)
}
