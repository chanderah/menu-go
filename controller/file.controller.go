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
		return
	}
	response.OK(c, files)
}

func UploadFile(c *gin.Context) {
	var req model.FileDetails
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := util.UploadFile(&req)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	response.OK(c, res)
}

func DeleteFile(c *gin.Context) {
	var req model.FileDetails
	c.ShouldBindJSON(&req)

	if err := util.RemoveFile(req.Dest); err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	response.Void(c)
}
