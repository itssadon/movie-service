package film

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const BaseUrl = "https://swapi.dev/api"

func GetAllFilms() (*FilmCollection, error) {
	result := &FilmCollection{}

	response, err := http.Get(fmt.Sprintf("%s%s", BaseUrl, "/films/"))
	if err != nil {
		return nil, err
	}

	body, readError := ioutil.ReadAll(response.Body)
	if readError != nil {
		return nil, readError
	}

	closeErr := response.Body.Close()
	if closeErr != nil {
		return nil, closeErr
	}

	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
