package core

//type RouterGroup struct {
//	Router       *gin.Engine      // 默认GIN路由引擎
//	PublicGroup  *gin.RouterGroup // 公开路由组
//	PrivateGroup *gin.RouterGroup // 权限路由组
//}
//
////
//// InitRouterGroup
////  @Description: 初始化路由组
////  @return *RouterGroup
////
//func InitRouterGroup() *RouterGroup {
//	// Gin路由组初始化
//	var router = gin.Default()
//	publicGroup := router.Group("")
//	privateGroup := router.Group("")
//	// 权限路由中间件配置
//	// 全局路由组初始化
//	return &RouterGroup{
//		Router:       router,
//		PublicGroup:  publicGroup,
//		PrivateGroup: privateGroup,
//	}
//}
