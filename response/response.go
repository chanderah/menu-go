package response

import (
	"github.com/chanderah/menu-go/util"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Void(c *gin.Context) {
	c.JSON(200, Response{
		Status:  200,
		Message: "success",
	})
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(200, Response{
		Status:  200,
		Message: "success",
		Data:    data,
	})
}

func Error(c *gin.Context, status int, message string) {
	if util.IsEmpty(status) {
		status = 500
	} else if util.IsEmpty(message) {
		message = "Something went wrong."
	}

	c.JSON(status, Response{
		Status:  status,
		Message: message,
	})
}
