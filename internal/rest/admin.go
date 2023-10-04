package rest

import (
	"github.com/gin-gonic/gin"
	"item-search/pkg/response"
)

type AdminRest struct {
}

func GetAdminRest() *AdminRest {
	return &AdminRest{}
}

func (ar *AdminRest) Info(c *gin.Context) {
	response.Send("OK", c)
}
