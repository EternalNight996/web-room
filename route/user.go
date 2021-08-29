// route/user.go
package route

import (
	"github.com/EternalNight996/web-room/controller"
	"github.com/EternalNight996/web-room/middleware"
	"github.com/kataras/iris/v12/core/router"
)

func routeUser(party router.Party) {
	party.Post("/login", controller.PostLogin)

	party.Post("/user", controller.PostUser)
	party.Get("/user", controller.GetUser)
	party.Put("/user", middleware.JWT.Serve, middleware.Logined, controller.PutUser)
}
