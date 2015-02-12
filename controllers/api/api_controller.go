package api_controllers

import "net/http"

type ApiController struct {
}

func (controller ApiController) Render(w http.ResponseWriter, r *http.Request, body string) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(body))
}
