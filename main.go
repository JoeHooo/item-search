package main

import (
	"item-search/internal/repository/db"
	"item-search/internal/repository/es"
	"item-search/internal/repository/redis"
	"item-search/internal/router"
	"item-search/pkg/config"
)

func main() {
	config.ReadConfig()
	db.Init()
	es.Init()
	redis.Init()
	router.Run()
}
