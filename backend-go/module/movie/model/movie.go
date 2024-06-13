package moviemodel

import (
	"cinema/common"
)

const EntityName = "Movie"
const TableName = "movies"

type Movie struct {
	common.Movie
}

func (Movie) TableName() string { return TableName }
