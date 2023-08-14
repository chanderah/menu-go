package controller

import (
	"net/http"

	"github.com/chanderah/menu-go/model"
	"github.com/chanderah/menu-go/response"
	"github.com/chanderah/menu-go/util"
	"github.com/gin-gonic/gin"
)

func GetFiles(c *gin.Context) {
	files, err := util.GetFiles()
	if err != nil {
		response.Error(c, 500, err.Error())
	}
	response.OK(c, files)
}

func UploadFile(c *gin.Context) {
	var req model.File
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
	}

	res, err := util.UploadFile("/user", "./")
	if err != nil {
		response.Error(c, 400, err.Error())
	}
	response.OK(c, res)
}

func DeleteFile(c *gin.Context) {
	var req model.File
	c.ShouldBindJSON(&req)

	if err := util.RemoveFile(req.File); err != nil {
		response.Error(c, 400, err.Error())
	}
	response.Void(c)
}
