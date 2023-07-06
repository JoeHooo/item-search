package repository

import (
	"context"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Ctx = context.Background()
	Rdb *redis.Client
	Db  *gorm.DB
	Es  *elasticsearch.Client
)
