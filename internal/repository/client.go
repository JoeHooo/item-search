package repository

import (
	"context"
	"github.com/elastic/go-elasticsearch/v7"
	gr "github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Ctx           = context.Background()
	Redis         *gr.Client
	MySQL         *gorm.DB
	Elasticsearch *elasticsearch.Client
	initialized   bool
)

func Init() {
	if initialized {
		return
	}
	InitElasticsearch()
	InitRedis()
	InitMySQL()
	initialized = true
}
