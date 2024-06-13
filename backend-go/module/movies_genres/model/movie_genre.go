package moviesgenresmodel

import "cinema/common"

const EntityName = "Movie Genre"
const TableName = "movies_genres"

type MovieGenre common.MovieGenre

func (MovieGenre) TableName() string { return TableName }
