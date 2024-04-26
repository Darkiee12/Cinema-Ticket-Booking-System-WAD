package moviemodel

import (
	"time"
)

const EntityName = "Movie"
const TableName = "movies"

type Movie struct {
	ImdbID        string    `gorm:"column:imdb_id"`
	Awards        string    `gorm:"column:awards"`
	Dvd           string    `gorm:"column:dvd"`
	ImdbRating    float64   `gorm:"column:imdb_rating"`
	ImdbVotes     int       `gorm:"column:imdb_votes"`
	Metascore     int       `gorm:"column:metascore"`
	OriginalTitle string    `gorm:"column:original_title"`
	Plot          string    `gorm:"column:plot"`
	Poster        string    `gorm:"column:poster"`
	Rated         string    `gorm:"column:rated"`
	Released      string    `gorm:"column:released"`
	Runtime       int       `gorm:"column:runtime"`
	Tagline       string    `gorm:"column:tagline"`
	Title         string    `gorm:"column:title"`
	Type          string    `gorm:"column:type"`
	Website       string    `gorm:"column:website;default:'N/A'"`
	Year          int       `gorm:"column:year"`
	Production    string    `gorm:"column:production"`
	TmdbID        int       `gorm:"column:tmdb_id"`
	BoxOffice     float64   `gorm:"column:box_office"`
	CreatedAt     time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
}

func (Movie) TableName() string { return TableName }
