// controller/index.go
package controller

import (
	"github.com/EternalNight996/web-room/database"
	"github.com/EternalNight996/web-room/service"
)

var db = database.DB
var messageService = service.NewMessage()
var userService = service.NewUser()
