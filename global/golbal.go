// Package global
// @Description: 全局实例
package global

import (
	"github.com/codycoding/goDuck/global/config"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	MysqlDb       *gorm.DB         // mysql数据库实例
	PostgresDb    *gorm.DB         // postgre数据库实例
	Redis         *redis.Client    // redis内存数据库实例
	Validator     *Validate        // 结构体验证器
	Vp            *viper.Viper     // 命令行处理实例
	Log           *zap.Logger      // 日志实例
	Config        config.Server    // 程序配置
	Route         *gin.Engine      // 程序路由组
	PublicRouter  *gin.RouterGroup // 公共路由
	PrivateRouter *gin.RouterGroup // 权限路由（已加载权限中间件）
)
