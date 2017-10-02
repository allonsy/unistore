// Main unistore runner file
package main

import (
	"fmt"
	t "github.com/allonsy/unistore/template"
	"log"
	"net/http"
)

func main() {
	err := t.ParseTemplates()
	if err != nil {
		log.Fatal(err)
	}
	templates := t.GetTemplates()
	fmt.Println(templates.DefinedTemplates())

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		err = templates.ExecuteTemplate(w, "html/home.html", nil)
		if err != nil {
			log.Panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
