package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/osondoar/vohtk/controllers"
	"github.com/osondoar/vohtk/controllers/api"
	"github.com/osondoar/vohtk/errors"
)

func main() {
	var mainController controllers.MainController
	var requestsController api_controllers.RequestsController

	InitDirectories()
	r := mux.NewRouter()
	r.HandleFunc("/", mainController.Index)
	r.Handle("/api/requests", apiHandler(requestsController.Post)).Methods("POST")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./assets/")))

	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}

// http://blog.golang.org/error-handling-and-go
type apiHandler func(http.ResponseWriter, *http.Request) *errors.AppError

func (fn apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		w.Header().Set("Content-Type", "application/json")
		jsonResponse, _ := json.Marshal(e)
		http.Error(w, string(jsonResponse), e.Code)
		log.Printf("%q. Code: %q", e.Message, e.Code)
	}
}

func InitDirectories() {
	err1 := os.MkdirAll("tmp/wav", 0775)
	err2 := os.MkdirAll("tmp/mfc", 0775)
	err3 := os.MkdirAll("tmp/mlf", 0775)

	if err1 != nil || err2 != nil || err3 != nil {
		log.Printf("Error creating directories: %q, %q, %q", err1, err2, err3)
	}
}
