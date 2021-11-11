package core

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/codycoding/goDuck/global"
)

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
		if syncedEnforcer, sErr := casbin.NewSyncedEnforcer("./resource/restful_keyMatch2.conf", adapter); sErr != nil {
			global.Log.Error(fmt.Sprintf("Casbin初始化失败:%s", sErr))
			return nil
		} else {
			if pErr := syncedEnforcer.LoadPolicy(); pErr != nil {
				global.Log.Error(fmt.Sprintf("Casbin读取策略失败:%s", pErr))
			} else {
				return syncedEnforcer
			}
		}
	}
	return nil
}
