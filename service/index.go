// service/index.go
package service

import (
	"github.com/EternalNight996/web-room/database"
)

// NewMessage get a message service
func NewMessage() MessageService {
	return MessageService{
		db: database.DB,
	}
}

// NewUser get a user service
func NewUser() UserService {
	return UserService{
		db: database.DB,
	}
}
