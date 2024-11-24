package core

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gvb_server/global"
)

func Redis() {
	redisCfg := global.Config.Redis
	sprintf := fmt.Sprintf("%s:%d", redisCfg.Host, redisCfg.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     sprintf,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	fmt.Println(sprintf)
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.Log.Error(err)
		//global.GVA_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.Log.Info("redis connect ping response:", pong)
		//global.GVA_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		//global.GVA_REDIS = client
	}
}
