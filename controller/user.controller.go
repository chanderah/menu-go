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
		response.Error(c, http.StatusUnprocessableEntity, "Failed to register the account.")
		return
	}
	response.Void(c)
}

func LoginUser(c *gin.Context) {
	var input, data model.User

	c.ShouldBindJSON(&input)
	if res := util.DB.First(&data, "username = ?", input.Username); res.Error != nil {
		response.Error(c, http.StatusNotFound, "Data not found!")
		return
	}
	if input.Password != data.Password {
		response.Error(c, http.StatusUnauthorized, "Bad Credentials!")
		return
	}
	if res := util.DB.Model(&data).Updates(input); res.Error != nil {
		response.Error(c, http.StatusBadRequest, res.Error.Error())
		return
	}
	response.Void(c)
}

func GetUsers(c *gin.Context) {
	var data []model.UserBasic
	if res := util.DB.Find(&[]model.User{}).Scan(&data); res.Error != nil {
		response.Error(c, http.StatusBadRequest, res.Error.Error())
		return
	}
	response.OK(c, data)
}

func FindUserById(c *gin.Context) {
	var data model.UserBasic
	c.ShouldBindJSON(&data)
	if res := util.DB.First(&model.User{}, "id = ?", data.Id).Scan(&data); res.Error != nil {
		response.Error(c, http.StatusNotFound, "Data not found!")
		return
	}
	response.OK(c, data)
}

func FindUserByUsername(c *gin.Context) {
	var data model.UserBasic
	c.ShouldBindJSON(&data)
	if res := util.DB.First(&model.User{}, "id = ?", data.Id).Scan(&data); res.Error != nil {
		response.Error(c, http.StatusNotFound, "Data not found!")
		return
	}
	response.OK(c, data)
}

func UpdateUser(c *gin.Context) {
	var input, data model.User
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if res := util.DB.First(&data, "id = ?", input.Id); res.Error != nil {
		response.Error(c, http.StatusNotFound, "Data not found!")
		return
	}
	if res := util.DB.Model(&data).Updates(input); res.Error != nil {
		response.Error(c, http.StatusBadRequest, res.Error.Error())
		return
	}
	response.Void(c)
}

func DeleteUser(c *gin.Context) {
	var data model.User

	c.ShouldBindJSON(&data)
	if res := util.DB.First(&data, "id = ?", data.Id); res.Error != nil {
		response.Error(c, http.StatusNotFound, "Data not found!")
		return
	}
	if res := util.DB.Delete(&data); res.Error != nil {
		response.Error(c, http.StatusBadRequest, res.Error.Error())
		return
	}
	response.Void(c)
}
