package controller

import (
	"net/http"

	"github.com/chanderah/menu-go/model"
	"github.com/chanderah/menu-go/response"
	"github.com/chanderah/menu-go/util"
	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var data model.User
	if err := c.ShouldBindJSON(&data); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if res := util.DB.Create(&data); res.Error != nil {
		response.AppError(c, res.Error.Error())
		return
	}
	response.Void(c)
}

func LoginUser(c *gin.Context) {
	var req, data model.User
	var user model.UserBasic
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if res := util.DB.First(&data, "username = ?", req.Username).Scan(&user); res.Error != nil {
		response.Error(c, http.StatusNotFound, "User not found!")
		return
	}
	if req.Password != data.Password {
		response.Error(c, http.StatusUnauthorized, "Bad Credentials!")
		return
	}
	if res := util.DB.Model(&data).Updates(req); res.Error != nil {
		response.AppError(c, res.Error.Error())
		return
	}
	response.OK(c, user)
}

func GetUsers(c *gin.Context) {
	var data []model.UserBasic
	if res := util.DB.Find(&model.User{}).Scan(&data); res.Error != nil {
		response.Error(c, http.StatusInternalServerError, res.Error.Error())
		return
	}
	response.OK(c, data)
}

func FindUserById(c *gin.Context) {
	var data model.UserBasic
	c.ShouldBindJSON(&data)

	if res := util.DB.First(&data, "id = ?", data.ID); res.Error != nil {
		response.Error(c, http.StatusNotFound, "Data not found!")
		return
	}
	response.OK(c, data)
}

func FindUserByUsername(c *gin.Context) {
	var data model.UserBasic
	c.ShouldBindJSON(&data)

	if res := util.DB.First(&data, "id = ?", data.ID); res.Error != nil {
		response.Error(c, http.StatusNotFound, "Data not found!")
		return
	}
	response.OK(c, data)
}

func UpdateUser(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if res := util.DB.First(&req, "id = ?", req.ID); res.Error != nil {
		response.Error(c, http.StatusNotFound, "Data not found!")
		return
	}
	if res := util.DB.Model(&req).Updates(req); res.Error != nil {
		response.AppError(c, res.Error.Error())
		return
	}
	response.Void(c)
}

func DeleteUser(c *gin.Context) {
	var req model.User
	c.ShouldBindJSON(&req)

	if res := util.DB.First(&req, "id = ?", req.ID); res.Error != nil {
		response.Error(c, http.StatusNotFound, "Data not found!")
		return
	}
	if res := util.DB.Delete(&req); res.Error != nil {
		response.AppError(c, res.Error.Error())
		return
	}
	response.Void(c)
}
