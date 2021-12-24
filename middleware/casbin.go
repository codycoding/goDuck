package middleware

import (
	"github.com/codycoding/goDuck/core"
	"github.com/codycoding/goDuck/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserToRoles 临时角色-用户对应表
type UserToRoles struct {
	UserId int64  `json:"userId" gorm:"account_myuser_id"`
	RoleId string `json:"roleId" gorm:"authority_role_authority_id"`
}

func (UserToRoles) TableName() string {
	return "authority_user_role"
}

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
		//obj := c.Request.URL.RequestURI()
		obj := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的所有角色
		userId := waitUse.UserInfo.AccountId
		var roleIds []string
		db := global.PostgresDb.Model(&UserToRoles{}).Select("authority_role_authority_id").Where("account_myuser_id = ?", userId)
		if err := db.Scan(&roleIds).Error; err != nil {
			core.FailWithMessage(http.StatusBadRequest, "权限角色查询失败", c)
			c.Abort()
			return
		}
		success := false
		// 循环角色判断策略中是否存在
		for _, roleId := range roleIds {
			// 获取用户的角色
			sub := roleId
			success, _ = global.Casbin.Enforce(sub, obj, act)
			if success {
				break
			}
		}
		// 开发环境跳过访问限制
		if global.Config.System.Env == "develop" || success {
			c.Next()
		} else {
			core.FailWithMessage(http.StatusBadRequest, "权限不足", c)
			c.Abort()
			return
		}
	}
}
