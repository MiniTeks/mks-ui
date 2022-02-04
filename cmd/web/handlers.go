package main

import (
	"html/template"
	"log"
	"net/http"
)

func (app *application) homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}

	tpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Template couldn't be parsed", 500)
		return
	}

	if tpl.Execute(w, app) != nil {
		log.Println(err.Error())
		http.Error(w, "Template couldn't be parsed", 500)
		return
	}
}
