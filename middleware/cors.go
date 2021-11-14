package middleware

import (
	"github.com/codycoding/goDuck/global"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// Cors
//  @Description: 处理跨域请求
//  @return gin.HandlerFunc
//
func Cors() gin.HandlerFunc {
	switch global.Config.Cors.CorsMode {
	case "none":
		// 关闭跨域
		global.Log.Info("cors中间已关闭")
		return func(c *gin.Context) {
			c.Next()
		}
	case "config":
		// 按配置处理跨域
		global.Log.Info("cors中间按自定义配置启动")
		return cors.New(cors.Config{
			AllowAllOrigins:        false,
			AllowOrigins:           global.Config.Cors.AllowOrigins,
			AllowOriginFunc:        nil,
			AllowMethods:           global.Config.Cors.AllowMethods,
			AllowHeaders:           global.Config.Cors.AllowHeaders,
			AllowCredentials:       global.Config.Cors.AllowCredentials,
			ExposeHeaders:          global.Config.Cors.ExposeHeaders,
			MaxAge:                 time.Duration(global.Config.Cors.MaxAge) * time.Hour,
			AllowWildcard:          true,
			AllowBrowserExtensions: true,
			AllowWebSockets:        true,
			AllowFiles:             true,
		})
	default:
		// 默认跨域设置
		global.Log.Info("cors中间默认配置启动")
		return cors.Default()
	}
}
