package controller

import (
	"github.com/chanderah/menu-go/response"
	"github.com/chanderah/menu-go/util"
	"github.com/gin-gonic/gin"
)

type Query struct {
	Query string `json:"query"`
}

func RunQuery(c *gin.Context) {
	var req = Query{}
	c.ShouldBindJSON(&req)

	result := []map[string]interface{}{}
	if err := util.DB.Raw(req.Query).Scan(&result); err != nil {
		response.Error(c, 400, err.Error.Error())
		return
	}
	response.OK(c, result)
}
