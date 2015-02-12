package controllers

import (
	"net/http"
)

type AppController struct {
}

func (controller AppController) Render(w http.ResponseWriter, r *http.Request, body string) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(body))
}
