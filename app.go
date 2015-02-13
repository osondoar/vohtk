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

var (
	logger, logFile = InitLogger()
)

func main() {
	var mainController controllers.MainController
	var requestsController api_controllers.RequestsController
	defer logFile.Close()
	InitDirectories()

	r := mux.NewRouter()
	r.Handle("/", htmlHandler(mainController.Index))
	r.Handle("/api/requests", apiHandler(requestsController.Post)).Methods("POST")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./assets/")))

	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}

// http://blog.golang.org/error-handling-and-go
type apiHandler func(http.ResponseWriter, *http.Request) *errors.AppError
type htmlHandler func(http.ResponseWriter, *http.Request)

func (fn apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		w.Header().Set("Content-Type", "application/json")
		jsonResponse, _ := json.Marshal(e)
		http.Error(w, string(jsonResponse), e.Code)
		log.Printf("%q. Code: %q", e.Message, e.Code)
	}
}

func (fn htmlHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.Printf("%q %q Referer: %q", r.UserAgent(), r.RemoteAddr, r.Referer())
	fn(w, r)
}

func InitDirectories() {
	err1 := os.MkdirAll("tmp/wav", 0775)
	err2 := os.MkdirAll("tmp/mfc", 0775)
	err3 := os.MkdirAll("tmp/mlf", 0775)

	if err1 != nil || err2 != nil || err3 != nil {
		log.Printf("Error creating directories: %q, %q, %q", err1, err2, err3)
	}
}

func InitLogger() (*log.Logger, *os.File) {
	f, err := os.OpenFile("logs/access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("error opening file: %q", err)
	}
	// defer f.Close()
	logger := log.New(f, "", log.Ldate|log.Ltime)
	return logger, f
}
