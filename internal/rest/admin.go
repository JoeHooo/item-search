package rest

import (
	"github.com/gin-gonic/gin"
	"item-search/pkg/response"
	"time"
)

type AdminRest struct {
}

func GetAdminRest() *AdminRest {
	return &AdminRest{}
}

// @BasePath /admin
//
// admin info
// @Summary admin info
// @Schemes
// @Description admin info
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {string} OK
// @Router /admin/info [get]
func (ar *AdminRest) Info(c *gin.Context) {
	time.Sleep(time.Millisecond * 100)
	response.Send("OK", c)
}
