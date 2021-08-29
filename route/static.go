// route/static.go
package route

import "github.com/kataras/iris/v12"

func routeStatic(app *iris.Application) {
	// 默认静态资源
	app.HandleDir("/", "./assets")
}
