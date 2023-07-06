package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Infof("Request URL: %s , start", c.Request.Host+c.Request.URL.Path)
		t := time.Now()
		c.Next()
		latency := time.Since(t)
		status := c.Writer.Status()
		log.Infof(
			"Request URL: %s , status: %d, Spend Time: %dms",
			c.Request.Host+c.Request.URL.Path,
			status,
			latency.Milliseconds(),
		)
	}
}
