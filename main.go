// Main unistore runner file
package main

import (
	"fmt"
	"github.com/allonsy/unistore/controller"
	t "github.com/allonsy/unistore/template"
	"log"
	"net/http"
)

func main() {
	err := t.ParseTemplates()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(t.GetTemplates().DefinedTemplates())
	controller.Init()
	http.ListenAndServe(":8080", nil)
}
