package core

import (
	"context"
	"github.com/codycoding/goDuck/global"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

//
// Redis
//  @Description: 获取redis连接实例
//  @return *redis.Client
//
func Redis() *redis.Client {
	redisCfg := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.Log.Error("Redis连接失败，错误:", zap.Any("err", err))
		return nil
	} else {
		global.Log.Info("Redis连接成功:", zap.String("pong", pong))
		return client
	}
}
