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

	result := map[string]interface{}{}
	util.DB.Raw(req.Query).Scan(&result)

	response.OK(c, result)
}
