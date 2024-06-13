package moviesgenresmodel

type Filter struct {
	ImdbID  string `json:"imdbID" form:"imdbID" query:"imdbID"`
	GenreID int    `json:"genreID" form:"genreID" query:"genreID"`
}
