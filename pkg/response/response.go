package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	MessageId int         `json:"message_id"`
	Message   string      `json:"message"`
	Result    bool        `json:"result"`
	Object    interface{} `json:"object"`
}

func Send(data interface{}, c *gin.Context) {
	c.JSON(
		http.StatusOK,
		Response{
			10000,
			"",
			true,
			data,
		},
	)
}

func Success(message string, data interface{}, c *gin.Context) {
	c.JSON(
		http.StatusOK,
		Response{
			10000,
			message,
			true,
			data,
		},
	)
}

func Failed(message string, data interface{}, c *gin.Context) {
	c.JSON(
		http.StatusInternalServerError,
		Response{
			99999,
			message,
			false,
			data,
		},
	)
}
