/* ---------------项目技术栈--------------
技术栈：
    Golang -> HTML -> CSS -> JavaScript
框架：
    后端: Iris -> 前端: React
    数据库ORM(Object Relational Mapping): xorm -> 多类型配置文件读取支持: Viper
数据存储：
    Mysql
技术:
    签发登录令牌：JWT
    异步请求后端数据: AJAX
-----------------------------------------*/

/* ---------------项目结构---------------
route 路由层: 负责将Client请求映射到对应的函数
middleware 中间层: 执行函数前后进行拦截处理
controller 控制层: 存放与Client请求对应的函数，根据请求调用业务层，并将数据进行格式封>
service 业务层: 调用持久层完成业务逻辑
dao 持久层: 理解sql执行到函数执行的封装，让ORM替代。
database 提供数据库链接: -
model 定义一系列的结构体： -
    dbe(数据库对应实体) 业务逻辑实体: 如User、Message
    reqo(request object) 请求数据实体: 对应controller中的每个方法
    repo(response object) 响应数据实体: 对应controller中的每个方法
config: 用的是viper读取配置，并提供单实例的配置文件实体供外访问
tool 工具层： -
assets 静态资源目录，存放静态资源（前端文件）图片
----------------------------------------*/

// web-room/index.go main
package main

import (
	"github.com/kataras/iris/v12"

	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"

	_ "github.com/go-sql-driver/mysql"

	"github.com/EternalNight996/web-room/config"
	"github.com/EternalNight996/web-room/middleware"
	"github.com/EternalNight996/web-room/route"
)

func main() {

	// 初始化iris
	app := iris.New()
	// 设置日志等级: disable fatal error warn info debug
	app.Logger().SetLevel(config.Viper.GetString("server.logger.level"))
	// 添加recover去恢复来自http-relative panics
	app.Use(recover.New())
	// 添加请求日志到终端
	app.Use(logger.New())
	// 允许全局option方法启动用CORS(Cross-Origin Resource Sharing)跨站资源共享
	app.AllowMethods(iris.MethodOptions)
	// 添加全局CORS跨站资源共享handler,导入中间件
	app.Use(middleware.CORS)
	// 路由层加载
	route.Route(app)
	//监听端口
	app.Run(iris.Addr(config.Viper.GetString("server.addr")), iris.WithoutServerError(iris.ErrServerClosed))
}
