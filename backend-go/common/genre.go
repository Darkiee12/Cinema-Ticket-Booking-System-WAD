package common

type Genre struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
}

func (Genre) TableName() string { return "genres" }

type MovieGenre struct {
	ImdbID  string   `gorm:"column:imdb_id" json:"imdbID"`
	Movies  []*Movie `gorm:"preload:false;foreignKey:ImdbID;references:ImdbID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"movies"`
	GenreID int      `gorm:"column:genre_id" json:"genreID"`
}

func (MovieGenre) TableName() string { return "movies_genres" }
