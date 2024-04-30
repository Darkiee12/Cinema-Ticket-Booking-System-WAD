package moviemodel

import (
	"time"
)

const EntityName = "Movie"
const TableName = "movies"

type Movie struct {
	ImdbID        string     `gorm:"column:imdb_id" json:"imdbID"`
	Awards        string     `gorm:"column:awards" json:"awards"`
	Dvd           string     `gorm:"column:dvd" json:"dvd"`
	ImdbRating    float64    `gorm:"column:imdb_rating" json:"imdbRating"`
	ImdbVotes     int        `gorm:"column:imdb_votes" json:"imdbVotes"`
	Metascore     int        `gorm:"column:metascore" json:"metascore"`
	OriginalTitle string     `gorm:"column:original_title" json:"originalTitle"`
	Plot          string     `gorm:"column:plot" json:"plot"`
	Poster        string     `gorm:"column:poster" json:"poster"`
	Rated         string     `gorm:"column:rated" json:"rated"`
	Released      string     `gorm:"column:released" json:"released"`
	Runtime       int        `gorm:"column:runtime" json:"runtime"`
	Tagline       string     `gorm:"column:tagline" json:"tagline"`
	Title         string     `gorm:"column:title" json:"title"`
	Type          string     `gorm:"column:type" json:"type"`
	Website       string     `gorm:"column:website;default:'N/A'" json:"website"`
	Year          int        `gorm:"column:year" json:"year"`
	Production    string     `gorm:"column:production" json:"production"`
	TmdbID        int        `gorm:"column:tmdb_id" json:"tmdbID"`
	BoxOffice     float64    `gorm:"column:box_office" json:"boxOffice"`
	CreatedAt     *time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (Movie) TableName() string { return TableName }
