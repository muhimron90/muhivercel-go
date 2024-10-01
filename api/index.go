package handler

// package main
import (
	"encoding/json"
	"net/http"
)

var app http.Handler

func renderJson(w http.ResponseWriter, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

type Greeting struct {
	Message string `json:"message"`
}

func greetHandFn(w http.ResponseWriter, r *http.Request) {
	data := []byte(`{"message": "Hello"}`)
	var greet Greeting
	err := json.Unmarshal(data, &greet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	renderJson(w, greet)
}

/*
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", greetHandFn)
	srv := &http.Server{
		Addr:    ":3030",
		Handler: mux,
	}
	srv.ListenAndServe()
}
*/

func init() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", greetHandFn)
	app = mux
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
