package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"item-search/internal/repository"
	"item-search/pkg/config"
)

func Init() {
	redisConf := config.Conf.Redis
	repository.Rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConf.Host, redisConf.Port),
		Password: redisConf.Password,
		DB:       redisConf.Db,
	})
}
