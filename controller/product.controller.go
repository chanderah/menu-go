package controller

import (
	"database/sql"
	"fmt"
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

func GetProducts(c *gin.Context) {
	var rowCount int64
	data := []model.Product{}

	paging := model.PagingInfo{}
	c.ShouldBindJSON(&paging)
	util.GetPaging(&paging)

	where := "name LIKE @v OR code LIKE @v OR CAST(price AS CHAR) LIKE @v"
	value := sql.Named("v", "%"+paging.Filter+"%")

	res := util.DB.Order(util.StringJoin(paging.SortField, paging.SortOrder)).Limit(paging.Limit).Offset(paging.Offset).Find(&data, where, value).Count(&rowCount)
	if res.Error != nil {
		response.AppError(c, res.Error.Error())
		return
	}
	response.Paging(c, data, rowCount)
}

func FindProductByCategory(c *gin.Context) {
	var rowCount int64
	data := []model.Product{}

	paging := model.PagingInfo{}
	c.ShouldBindJSON(&paging)
	util.GetPaging(&paging)

	where := fmt.Sprintf("%s = '%s' AND name LIKE '%%%[3]s%%' OR code LIKE '%%%[3]s%%' OR CAST(price AS CHAR) LIKE '%%%[3]s%%'", paging.Field.Column, paging.Field.Value, paging.Filter)
	page := util.DB.Limit(paging.Limit).Offset(paging.Offset)
	res := page.Order(paging.SortField+" "+paging.SortOrder).Find(&data, where).Count(&rowCount)
	if res.Error != nil {
		response.AppError(c, res.Error.Error())
		return
	}
	response.Paging(c, data, rowCount)
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
