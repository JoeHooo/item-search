package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"item-search/docs"
	middleware2 "item-search/internal/middleware"
	"item-search/internal/rest"
	"item-search/pkg/config"
	"net/http"
	"time"
)

func Run() {
	router := gin.Default()
	router.Use(middleware2.Cors(), middleware2.Logger())

	serverConf := config.Conf.Server
	base := router.Group(serverConf.ContextPath)
	admin := base.Group("/admin")
	{
		adminRest := rest.GetAdminRest()
		admin.GET("/info", adminRest.Info)
	}

	initSwagger(router, serverConf.ContextPath)

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

func initSwagger(router *gin.Engine, basePath string) {
	docs.SwaggerInfo.BasePath = basePath
	router.GET(
		fmt.Sprintf("%s/swagger/*any", basePath),
		ginSwagger.WrapHandler(
			swaggerFiles.Handler,
			ginSwagger.URL(fmt.Sprintf("http://127.0.0.1:8080%s/swagger/doc.json", basePath)),
			ginSwagger.DefaultModelsExpandDepth(-1),
		),
	)
}
