// route/token.go
package route

import (
	"github.com/EternalNight996/web-room/controller"
	"github.com/EternalNight996/web-room/middleware"
	"github.com/kataras/iris/v12/core/router"
)

func routeToken(party router.Party) {
	party.Get("/token/info", middleware.JWT.Serve, middleware.Logined, controller.GetTokenInfo)
}
