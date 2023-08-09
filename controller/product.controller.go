package controller

import (
	"database/sql"
	"net/http"

	"github.com/chanderah/menu-go/model"
	"github.com/chanderah/menu-go/response"
	"github.com/chanderah/menu-go/util"
	"github.com/gin-gonic/gin"
)

type ProductsPaging struct {
	model.Product
	model.PagingInfo
}

// func GetProducts(c *gin.Context) {
// 	var data = []model.Product{}
// 	if res := util.DB.Find(&data); res.Error != nil {
// 		response.Error(c, http.StatusInternalServerError, res.Error.Error())
// 		return
// 	}
// 	response.OK(c, data)
// }

func GetProducts(c *gin.Context) {
	paging := model.PagingInfo{}
	c.ShouldBindJSON(&paging)
	util.GetPaging(&paging)

	where := "name LIKE @v OR code LIKE @v OR CAST(price AS CHAR) LIKE @v"
	value := sql.Named("v", "%"+paging.Filter+"%")

	var rowCount int64
	data := []model.Product{}
	res := util.DB.Order(util.StringJoin(paging.SortField, paging.SortOrder)).Limit(paging.Limit).Offset(paging.Offset).Find(&data, where, value).Count(&rowCount)
	if res.Error != nil {
		response.AppError(c, res.Error.Error())
		return
	}
	response.OK(c, data)
}

func FindProductById(c *gin.Context) {
	var data = model.Product{}
	c.ShouldBindJSON(&data)

	if res := util.DB.First(&data, "id = ?", data.ID); res.Error != nil {
		response.Error(c, http.StatusNotFound, "Data not found!")
		return
	}
	response.OK(c, data)
}

func FindProductByCategory(c *gin.Context) {
	var req model.Product
	c.ShouldBindJSON(&req)

	data := []model.Product{}
	if res := util.DB.Find(&data, "category = ?", req.Category); res.Error != nil {
		response.AppError(c, res.Error.Error())
		return
	}
	response.OK(c, data)
}

func CreateProduct(c *gin.Context) {
	var req model.Product
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if res := util.DB.Create(&req); res.Error != nil {
		response.AppError(c, res.Error.Error())
		return
	}
	response.Void(c)
}

func UpdateProduct(c *gin.Context) {
	var req model.Product
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

func DeleteProduct(c *gin.Context) {
	var data model.Product

	c.ShouldBindJSON(&data)
	if res := util.DB.First(&data, "id = ?", data.ID); res.Error != nil {
		response.Error(c, http.StatusNotFound, "Data not found!")
		return
	}
	if res := util.DB.Delete(&data); res.Error != nil {
		response.AppError(c, res.Error.Error())
		return
	}
	response.Void(c)
}
