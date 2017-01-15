package go_couchpotato

import (
	"github.com/juju/errors"
)

func (cc *CouchpotatoClient) SearchMovies(term string) ([]Movie, error) {

	if term == "" {
		return nil, errors.New("No term specified")
	}

	res := &SearchResponse{}

	err := cc.DoRequest("GET", "search", map[string]string{"q":term}, nil, res)

	if err != nil {
		return nil, errors.Annotate(err, "Failed to lookup movie")
	}

	if res.Success != true {
		return nil, errors.Errorf("Movie lookup returned failure")
	}

	return res.Movies, nil
}

func (cc *CouchpotatoClient) AddMovie(title string, imdbid string) (Movie, error) {
	res := &AddResponse{}

	err := cc.DoRequest("GET", "movie.add", map[string]string{"title":title, "identifier":imdbid}, nil, res)

	if err != nil {
		return Movie{}, errors.Annotate(err, "Failed to add movie")
	}

	if res.Success != true {
		return Movie{}, errors.Errorf("Movie add returned failure")
	}

	return res.Movie, nil
}