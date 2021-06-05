package films

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
	Make HTTP Get request to swapi.dev to get movile list
*/
func getFilmList() ([]Film, error) {
	// make the request
	res, err := http.Get("https://swapi.dev/api/films/")

	// check for response error
	if err != nil {
		return nil, err
	}

	// read all the response body
	data, _ := ioutil.ReadAll(res.Body)

	// close the response body
	res.Body.Close()

	fmt.Printf("%s\n", data)

	isValid := json.Valid(data)
	fmt.Println(isValid)

	// Structure response
	films := make([]Film, 0)

	return films, nil
}
