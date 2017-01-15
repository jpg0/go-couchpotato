package go_couchpotato

type SearchResponse struct {
	Movies []Movie `json:"movies"`
	Success bool `json:"success"`
}

type Movie struct {
	Rating struct {
		       Imdb []int `json:"imdb"`
	       } `json:"rating,omitempty"`
	TmdbID int `json:"tmdb_id"`
	ActorRoles struct {
	       } `json:"actor_roles"`
	ViaImdb bool `json:"via_imdb,omitempty"`
	ViaTmdb bool `json:"via_tmdb"`
	Titles []string `json:"titles"`
	Imdb string `json:"imdb"`
	Year int `json:"year"`
	Images struct {
		       PosterOriginal []interface{} `json:"poster_original"`
		       Poster []string `json:"poster"`
		       Actors struct {
				      } `json:"actors"`
		       BackdropOriginal []interface{} `json:"backdrop_original"`
		       Backdrop []interface{} `json:"backdrop"`
	       } `json:"images"`
	Plot string `json:"plot"`
	Genres []string `json:"genres"`
	InLibrary bool `json:"in_library"`
	Released string `json:"released,omitempty"`
	OriginalTitle string `json:"original_title"`
	Directors []string `json:"directors,omitempty"`
	Writers []interface{} `json:"writers,omitempty"`
	InWanted bool `json:"in_wanted"`
	Actors []string `json:"actors,omitempty"`
	Runtime int `json:"runtime,omitempty"`
	Type string `json:"type"`
	Tagline string `json:"tagline,omitempty"`
	Mpaa string `json:"mpaa,omitempty"`
}