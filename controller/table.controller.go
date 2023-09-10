package controller

import (
	"fmt"
	"net/http"

	"github.com/chanderah/menu-go/model"
	"github.com/chanderah/menu-go/response"
	"github.com/chanderah/menu-go/util"
	"github.com/gin-gonic/gin"
)

func GetTables(c *gin.Context) {
	var rowCount int64
	data := []model.Table{}

	paging := model.PagingInfo{}
	c.ShouldBindJSON(&paging)
	util.GetPaging(&paging)

	filter := "1=1"
	if !util.IsEmpty(paging.Filter) {
		filter = fmt.Sprintf("name LIKE '%%%s%%'", paging.Filter)
	}

	res := util.DB.Model(&model.Table{}).Where(filter).Count(&rowCount).
		Order(util.StringJoin(paging.SortField, paging.SortOrder)).Limit(paging.Limit).Offset(paging.Offset).
		Find(&data)
	if res.Error != nil {
		// if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		response.AppError(c, res.Error.Error())
		return
		// }
	}
	response.Paging(c, data, rowCount)
}

func FindTableById(c *gin.Context) {
	var data = model.Table{}
	c.ShouldBindJSON(&data)

	if res := util.DB.First(&data, "id = ?", data.ID); res.Error != nil {
		response.Error(c, http.StatusNotFound, "Data not found!")
		return
	}
	response.OK(c, data)
}

func CreateTable(c *gin.Context) {
	var req model.Table
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

func UpdateTable(c *gin.Context) {
	var req = model.Table{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var data = model.Table{}
	if res := util.DB.First(&data, "id = ?", req.ID); res.Error != nil {
		response.Error(c, http.StatusNotFound, "Data not found!")
		return
	}

	if res := util.DB.Model(&req).Updates(&req); res.Error != nil {
		response.AppError(c, res.Error.Error())
		return
	}
	response.Void(c)
}

func DeleteTable(c *gin.Context) {
	var req = model.Table{}
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
