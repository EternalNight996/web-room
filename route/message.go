// route/message.go
package route

import (
	"github.com/EternalNight996/web-room/controller"
	"github.com/EternalNight996/web-room/middleware"
	"github.com/kataras/iris/v12/core/router"
)

func routeMessage(party router.Party) {
	party.Post("/message", middleware.JWT.Serve, middleware.Logined, controller.PostMessage)
	party.Get("/message", middleware.JWT.Serve, middleware.Logined, controller.GetMessage)
}
