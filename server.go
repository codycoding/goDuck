package framework

import (
	"fmt"
	"github.com/codycoding/mini-go-app/core"
	"github.com/codycoding/mini-go-app/global"
	"github.com/codycoding/mini-go-app/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//
// InitialApp
//  @Description: 程序初始化
//
func InitialApp() {
	global.Vp = core.Viper()                                  // 配置读取：命令行参数|默认文件
	global.Log = core.Zap()                                   // 初始化日志库
	global.MysqlDb = core.GormMysql(global.Config.SlpDataDb)  // 初始化Mysql数据库
	global.PostgresDb = core.GormPostgres(global.Config.DgDb) // 初始化Postgres数据库
	global.Redis = core.Redis()                               // 初始化Redis连接
	// 路由定义
	// 初始化默认路由组
	global.Route = gin.Default()
	// 跨域中间件
	global.Route.Use(middleware.Cors())
	//
	global.PublicRouter = global.Route.Group("")  // 无权限路由组
	global.PrivateRouter = global.Route.Group("") // 权限路由组
	// 权限中间件
	global.PrivateRouter.Use(middleware.JwtAuth())
}

//
// RunApp
//  @Description: 程序运行方法
//
func RunApp() {
	// Gin服务开启
	address := fmt.Sprintf(":%d", global.Config.System.Addr)
	server := core.InitServer(address, global.Route)
	global.Log.Info("服务运行端口: ", zap.String("address", address))
	global.Log.Error(server.ListenAndServe().Error())
}
