package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/itssadon/movie-service/film"
)

const serviceBasePath = "/"

type ResponseObj struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		response := ResponseObj{true, "Welcome to our Movie Listing Service"}

		js, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(js)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	http.HandleFunc(serviceBasePath, baseHandler)
	film.SetupRoutes(serviceBasePath)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5300"
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
