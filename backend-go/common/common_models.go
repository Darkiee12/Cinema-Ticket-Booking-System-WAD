package common

import "time"

type Auditorium struct {
	SQLModel `json:",inline"`
	Name     string        `json:"name" gorm:"column:name;"`
	Seats    int           `json:"seats" gorm:"column:seats;"`
	CinemaID int           `json:"-" gorm:"column:cinema_id;"`
	Cinema   *SimpleCinema `json:"cinema" gorm:"preload:false;foreignKey:CinemaID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Auditorium) TableName() string { return "auditoriums" }

func (c *Auditorium) Mask(isAdminOrOwner bool) {
	c.GenUID(DbTypeAuditorium)
	if cinema := c.Cinema; cinema != nil {
		cinema.Mask(isAdminOrOwner)
	}
}

type Movie struct {
	ImdbID        string     `gorm:"column:imdb_id;primary_key" json:"imdbID"`
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
	Genres        []*Genre   `gorm:"many2many:movies_genres;joinForeignKey:ImdbID;joinReferences:GenreID" json:"genres"`
}

func (Movie) TableName() string { return "movies" }

type SimpleMovie struct {
	ImdbID        string  `gorm:"column:imdb_id;primary_key" json:"imdbID"`
	OriginalTitle string  `gorm:"column:original_title" json:"originalTitle"`
	Plot          string  `gorm:"column:plot" json:"plot"`
	Runtime       int     `gorm:"column:runtime" json:"runtime"`
	Title         string  `gorm:"column:title" json:"title"`
	Year          int     `gorm:"column:year" json:"year"`
	Poster        string  `gorm:"column:poster" json:"poster"`
	ImdbRating    float64 `gorm:"column:imdb_rating" json:"imdbRating"`
}

func (SimpleMovie) TableName() string { return "movies" }

type Cinema struct {
	SQLModel    `json:",inline"`
	OwnerID     int         `json:"-" gorm:"column:owner_id;"`
	Owner       *SimpleUser `json:"owner" gorm:"preload:false;foreignKey:OwnerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name        string      `json:"name" gorm:"column:name;"`
	Address     string      `json:"address"  gorm:"column:address;"`
	Capacity    int         `json:"capacity" gorm:"column:capacity;"`
	Email       string      `json:"email" gorm:"column:email;"`
	PhoneNumber string      `json:"phone_number" gorm:"column:phone_number;"`
}

func (Cinema) TableName() string {
	return "cinemas"
}

func (u *Cinema) Mask(isAdmin bool) {
	u.GenUID(DbTypeCinema)
}

type SimpleCinema struct {
	SQLModel `json:",inline"`
	OwnerID  int    `json:"-" gorm:"column:owner_id;"`
	Name     string `json:"name" gorm:"column:name;"`
}

func (SimpleCinema) TableName() string {
	return "cinemas"
}

func (u *SimpleCinema) Mask(isAdmin bool) {
	u.GenUID(DbTypeCinema)
}

type SimpleUser struct {
	SQLModel `json:",inline"`
	Name     string `json:"name" gorm:"column:name;"`
	Email    string `json:"email" gorm:"column:email;"`
	Tier     string `json:"tier" gorm:"column:tier;"`
}

func (SimpleUser) TableName() string {
	return "users"
}

func (u *SimpleUser) Mask(isAdmin bool) {
	u.GenUID(DbTypeUser)
}

type Show struct {
	ID           int          `json:"id" gorm:"column:id;primary_key"`
	Date         *Date        `json:"date" gorm:"column:date" `
	StartTime    *Time        `json:"startTime" gorm:"column:start_time"`
	EndTime      *Time        `json:"endTime" gorm:"column:end_time"`
	AuditoriumID int64        `json:"-" gorm:"column:auditorium_id"`
	Auditorium   *Auditorium  `json:"auditorium,omitempty" gorm:"preload:false;foreignKey:AuditoriumID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ImdbID       string       `json:"imdbID" gorm:"column:imdb_id"`
	Movie        *SimpleMovie `json:"movie,omitempty" gorm:"preload:false;foreignKey:ImdbID;references:ImdbID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt    *time.Time   `json:"-" gorm:"column:created_at"`
	UpdatedAt    *time.Time   `json:"-" gorm:"column:updated_at"`
}

func (Show) TableName() string { return "shows" }
