package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	middleware2 "item-search/internal/middleware"
	"item-search/internal/rest"
	"item-search/pkg/config"
	"net/http"
	"time"
)

func Run() {
	router := gin.Default()
	router.Use(
		middleware2.Cors(),
		middleware2.Logger(),
	)

	serverConf := config.Config.Server
	base := router.Group(serverConf.ContextPath)
	admin := base.Group("/admin")
	{
		adminRest := rest.GetAdminRest()
		admin.GET("/info", adminRest.Info)
	}

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", serverConf.Port),
		Handler:        router,
		ReadTimeout:    time.Duration(serverConf.ReadTimeOut) * time.Millisecond,
		WriteTimeout:   time.Duration(serverConf.WriteTimeOut) * time.Millisecond,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("Server start error, %s", err)
		return
	}
}
