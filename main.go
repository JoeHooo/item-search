package main

import (
	"item-search/internal/repository"
	"item-search/internal/router"
	"item-search/pkg/config"
)

func main() {
	config.ReadConfig()
	repository.Init()
	router.Run()
}
