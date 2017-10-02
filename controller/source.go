package controller

import (
	t "github.com/allonsy/unistore/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func handleSource(w http.ResponseWriter, req *http.Request) {
	fname := req.URL.Path[1:]
	err := t.ExecuteTemplate(w, fname, nil)
	if err != nil {
		log.Println(err)
	}
}

func handleJS(w http.ResponseWriter, req *http.Request) {
	fname := filepath.Join("template", filepath.FromSlash(req.URL.Path[1:]))
	contents, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Println(err)
	}
	w.Write(contents)
}
