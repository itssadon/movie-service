package films

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const filmsPath = "movies"

func handleFilmList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		filmList, err := getFilmList()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		newList := make([]string, 0)
		for _, s := range filmList {
			newList = append(newList, s.Title)
		}

		filmsJson, err := json.Marshal(newList)
		if err != nil {
			log.Fatal(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(filmsJson)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func SetupRoutes(serviceBasePath string) {
	filmListHandler := http.HandlerFunc(handleFilmList)

	http.Handle(fmt.Sprintf("%s%s", serviceBasePath, filmsPath), filmListHandler)
}
