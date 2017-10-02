package controller

import (
	t "github.com/allonsy/unistore/template"
	"net/http"
)

func getHome(w http.ResponseWriter, req *http.Request) {
	t.ExecuteTemplate(w, "html/home.html", nil)
}
