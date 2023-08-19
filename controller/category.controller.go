package controller

import (
	"net/http"

	"github.com/chanderah/menu-go/model"
	"github.com/chanderah/menu-go/response"
	"github.com/chanderah/menu-go/util"
	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	data := []model.Category{}
	if res := util.DB.Find(&data); res.Error != nil {
		response.AppError(c, res.Error.Error())
		return
	}
	response.OK(c, &data)
}

func FindCategoryById(c *gin.Context) {
	var data = model.Category{}
	c.ShouldBindJSON(&data)

	if res := util.DB.First(&data, "id = ?", data.ID); res.Error != nil {
		response.Error(c, http.StatusNotFound, "Data not found!")
		return
	}
	response.OK(c, data)
}

func CreateCategory(c *gin.Context) {
	var req model.Category
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	if res := util.DB.Create(&req); res.Error != nil {
		response.Error(c, http.StatusUnprocessableEntity, res.Error.Error())
		return
	}
	response.Void(c)
}

func UpdateCategory(c *gin.Context) {
	var req model.Category
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if res := util.DB.First(&model.Category{}, "id = ?", req.ID); res.Error != nil {
		response.Error(c, http.StatusNotFound, "Data not found!")
		return
	}
	if res := util.DB.Model(&req).Updates(&req); res.Error != nil {
		response.AppError(c, res.Error.Error())
		return
	}
	response.Void(c)
}

func DeleteCategory(c *gin.Context) {
	var req model.Category
	c.ShouldBindJSON(&req)

	if res := util.DB.First(&req, "id = ?", req.ID); res.Error != nil {
		response.Error(c, http.StatusNotFound, "Data not found!")
		return
	}
	if res := util.DB.Delete(&req); res.Error != nil {
		response.AppError(c, res.Error.Error())
		return
	}
	// set product with that ID to null
	res := util.DB.Model(&model.Product{}).Where("category_id = ?", req.ID).Update("category_id", nil)
	if res.Error != nil {
		response.AppError(c, res.Error.Error())
		return
	}
	response.Void(c)
}
