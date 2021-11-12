package middleware

import (
	"github.com/codycoding/goDuck/global"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// CorsMiddleware
//  @Description: 处理跨域请求
//  @return gin.HandlerFunc
//
func CorsMiddleware() gin.HandlerFunc {
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
}
