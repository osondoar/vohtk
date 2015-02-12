package controllers

import (
	"net/http"
	"path"
	"text/template"
)

type MainController struct {
	AppController
}

func (controller MainController) Index(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("templates", "index.html")
	template, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := template.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
