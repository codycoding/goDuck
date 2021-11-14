package middleware

import (
	"fmt"
	"github.com/codycoding/goDuck/core"
	"github.com/codycoding/goDuck/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

//
// CasbinHandler
//  @Description: API访问权限Casbin鉴权
//  @return gin.HandlerFunc
//
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 正常鉴权
		claims, _ := c.Get("claims")
		waitUse := claims.(*CustomClaims)
		// 获取请求的URI
		obj := fmt.Sprintf("%s%s", global.Config.System.UrlPreFix, c.Request.URL.RequestURI())
		fmt.Println(obj)
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := waitUse.UserInfo.RoleId
		// 判断策略中是否存在
		success, _ := global.Casbin.Enforce(sub, obj, act)
		// 开发环境 或 超管角色跳过访问限制
		if global.Config.System.Env == "develop" || sub == "9999" || success {
			c.Next()
		} else {
			core.FailWithMessage(http.StatusBadRequest, "权限不足", c)
			c.Abort()
			return
		}
	}
}
