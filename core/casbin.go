package core

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/codycoding/goDuck/global"
)

var casbinRestfulModel = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && keyMatch2(r.obj, p.obj) && regexMatch(r.act, p.act)
`

//
// InitCasbin
//  @Description: Casbin实例初始化
//  @return *casbin.SyncedEnforcer
//
func InitCasbin() *casbin.SyncedEnforcer {
	if adapter, err := gormadapter.NewAdapterByDB(global.PostgresDb); err != nil {
		global.Log.Error(fmt.Sprintf("Casbin数据库适配器初始化失败:%s", err))
		return nil
	} else {
		// 字符串模型转换
		casbinModel, _ := model.NewModelFromString(casbinRestfulModel)
		if syncedEnforcer, sErr := casbin.NewSyncedEnforcer(casbinModel, adapter); sErr != nil {
			global.Log.Error(fmt.Sprintf("Casbin初始化失败:%s", sErr))
			return nil
		} else {
			if pErr := syncedEnforcer.LoadPolicy(); pErr != nil {
				global.Log.Error(fmt.Sprintf("Casbin读取策略失败:%s", pErr))
			} else {
				global.Log.Info("Casbin实例初始化成功")
				return syncedEnforcer
			}
		}
	}
	return nil
}
