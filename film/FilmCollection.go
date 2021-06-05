package film

import (
	"github.com/itssadon/movie-service/collection"
	"github.com/itssadon/movie-service/errors"
)

type FilmCollection struct {
	errors.Error

	collection.SWAPIResultCollection
	Results []Film `json:"results"`
}
