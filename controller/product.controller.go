package controller

import (
	"fmt"
	"net/http"

	"github.com/chanderah/menu-go/model"
	"github.com/chanderah/menu-go/response"
	"github.com/chanderah/menu-go/util"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var rowCount int64
	data := []model.Product{}

	paging := model.PagingInfo{}
	c.ShouldBindJSON(&paging)
	util.GetPaging(&paging)

	where := "1=1"
	if !util.IsEmpty(paging.Filter) {
		where = fmt.Sprintf("name LIKE '%%%[1]s%%' OR code LIKE '%%%[1]s%%' OR CAST(price AS CHAR) LIKE '%%%[1]s%%'", paging.Filter)
	}
	res := util.DB.Order(util.StringJoin(paging.SortField, paging.SortOrder)).Limit(paging.Limit).Offset(paging.Offset).Find(&data, where).Count(&rowCount)
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

	where := fmt.Sprintf("categoryId = %d AND name LIKE '%%%[2]s%%' OR code LIKE '%%%[2]s%%' OR CAST(price AS CHAR) LIKE '%%%[2]s%%'", paging.Field.Value, paging.Filter)
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

	if !util.IsEmpty(req.Image) {
		imageUrl, err := util.UploadFile(&model.FileDetails{
			Dest: "/image.jpg",
			File: req.Image,
		})
		if err != nil {
			response.Error(c, 400, err.Error())
			return
		}
		req.Image = imageUrl
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

	var data = model.Product{}
	if res := util.DB.First(&data, "id = ?", req.ID); res.Error != nil {
		response.Error(c, http.StatusNotFound, "Data not found!")
		return
	}

	if req.Image != data.Image {
		// uploadImg
		imageUrl, err := util.UploadFile(&model.FileDetails{
			Dest: "/image.jpg",
			File: req.Image,
		})
		if err != nil {
			response.Error(c, 500, err.Error())
			return
		}
		req.Image = imageUrl

		if err := util.RemoveFile(data.Image); err != nil {
			response.Error(c, 500, err.Error())
			return
		}
	}

	fmt.Sprintln(req)
	if res := util.DB.Model(&req).Updates(req); res.Error != nil {
		response.AppError(c, res.Error.Error())
		return
	}
	response.Void(c)
}

func DeleteProduct(c *gin.Context) {
	var req model.Product

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
