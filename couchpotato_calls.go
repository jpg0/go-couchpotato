package go_couchpotato

import (
	"github.com/juju/errors"
)

func (cc *CouchpotatoClient) MovieSeachLookup(term string) ([]Movie, error) {

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