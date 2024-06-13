package genremodel

import "cinema/common"

const EntityName = "Genre"
const TableName = "genres"

type Genre common.Genre

func (Genre) TableName() string { return TableName }
