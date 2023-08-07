package controller

import (
	"net/http"

	"github.com/chanderah/menu-go/model"
	"github.com/chanderah/menu-go/response"
	"github.com/chanderah/menu-go/util"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var data []model.Product
	if res := util.DB.Find(&[]model.Product{}).Scan(&data); res.Error != nil {
		response.Error(c, http.StatusBadRequest, res.Error.Error())
		return
	}
	response.OK(c, data)
}

func FindProductById(c *gin.Context) {
	var data model.UserBasic
	c.ShouldBindJSON(&data)
	if res := util.DB.First(&model.User{}, "id = ?", data.ID).Scan(&data); res.Error != nil {
		response.Error(c, http.StatusNotFound, "Data not found!")
		return
	}
	response.OK(c, data)
}

func FindProductByCategory(c *gin.Context) {
	var req model.Product
	var data []model.Product
	c.ShouldBindJSON(&req)
	if res := util.DB.Where(&model.Product{Category: req.Category}).Find(&[]model.Product{}).Scan(&data); res.Error != nil {
		response.Error(c, http.StatusNotFound, "Data not found!")
		return
	}
	response.OK(c, req)
}

func CreateProduct(c *gin.Context) {
	var req model.Product
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if res := util.DB.Create(&req); res.Error != nil {
		response.Error(c, http.StatusBadRequest, "Failed to create the product.")
		return
	}
	response.Void(c)
}

func UpdateProduct(c *gin.Context) {
	var input, data model.User
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if res := util.DB.First(&data, "id = ?", input.ID); res.Error != nil {
		response.Error(c, http.StatusNotFound, "Data not found!")
		return
	}
	if res := util.DB.Model(&data).Updates(input); res.Error != nil {
		response.Error(c, http.StatusBadRequest, res.Error.Error())
		return
	}
	response.Void(c)
}

func DeleteProduct(c *gin.Context) {
	var data model.User

	c.ShouldBindJSON(&data)
	if res := util.DB.First(&data, "id = ?", data.ID); res.Error != nil {
		response.Error(c, http.StatusNotFound, "Data not found!")
		return
	}
	if res := util.DB.Delete(&data); res.Error != nil {
		response.Error(c, http.StatusBadRequest, res.Error.Error())
		return
	}
	response.Void(c)
}
