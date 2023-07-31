package repository

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"item-search/pkg/config"
)

func InitRedis() {
	redisConf := config.Config.Redis
	Rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConf.Host, redisConf.Port),
		Password: redisConf.Password,
		DB:       redisConf.Db,
	})
}
