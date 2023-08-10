package response

import (
	"net/http"

	"github.com/chanderah/menu-go/util"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status   int         `json:"status"`
	Message  string      `json:"message"`
	RowCount int64       `json:"rowCount,omitempty"`
	Data     interface{} `json:"data,omitempty"`
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

func Paging(c *gin.Context, data interface{}, rowCount int64) {
	c.JSON(200, Response{
		Status:   200,
		Message:  "success",
		Data:     data,
		RowCount: rowCount,
	})
}

func Error(c *gin.Context, status int, message string) {
	c.JSON(status, Response{
		Status:  status,
		Message: message,
	})
}

func AppError(c *gin.Context, message string) {
	if util.IsEmpty(message) {
		message = "Something went wrong."
	}
	Error(c, http.StatusInternalServerError, message)
}
