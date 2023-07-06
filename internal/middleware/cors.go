package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	origin                        = "Origin"
	accessControlAllowOrigin      = "Access-Control-Allow-Origin"
	accessControlAllowMethods     = "Access-Control-Allow-Methods"
	accessControlAllowHeaders     = "Access-Control-Allow-Headers"
	accessControlAllowCredentials = "Access-Control-Allow-Credentials"
	accessControlExposeHeaders    = "Access-Control-Expose-Methods"
	accessControlMaxAge           = "Access-Control-Max-Age"
	contentType                   = "content-type"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get(origin)
		c.Header(accessControlAllowOrigin, origin)
		c.Header(accessControlAllowHeaders, "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
		c.Header(accessControlAllowMethods, "POST,GET,OPTIONS,DELETE,PUT")
		c.Header(accessControlExposeHeaders, "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header(accessControlAllowCredentials, "true")
		c.Header(accessControlMaxAge, "172800")
		c.Header(contentType, "application/json")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
