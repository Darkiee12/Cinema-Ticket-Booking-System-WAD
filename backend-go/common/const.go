package common

import (
	"errors"
	"log"
)

const (
	DbTypeCinema = iota + 1
	DbTypeUser
	DbTypeAuditorium
	DbTypeCompany
)

const (
	CurrentUser = "user"
)

const (
	TopicUserLikeRestaurant    = "TopicUserLikeRestaurant"
	TopicUserDislikeRestaurant = "TopicUserDislikeRestaurant"
)

var (
	RecordNotFound = errors.New("record not found")
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery error", err)
	}
}
