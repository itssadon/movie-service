package film

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const moviesPath = "movies"

func handleFilmsList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		filmsList, err := GetAllFilms()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		moviesJson, err := json.Marshal(filmsList)
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(moviesJson)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func SetupRoutes(serviceBasePath string) {
	movieListHandler := http.HandlerFunc(handleFilmsList)

	http.Handle(fmt.Sprintf("%s%s", serviceBasePath, moviesPath), movieListHandler)
}
