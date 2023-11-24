package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/chanderah/menu-go/model"
	"github.com/chanderah/menu-go/response"
	"github.com/chanderah/menu-go/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetLiveOrders(c *gin.Context) {
	data := []model.Order{}

	req := model.GetLiveOrder{}
	c.ShouldBindJSON(&req)

 	res := util.DB.Raw(fmt.Sprintf(`CALL USP_GET_LIVE_ORDERS(%d, %d)`, req.ID, req.Limit)).Scan(&data);
	if res.Error != nil {
		response.AppError(c, res.Error.Error())
		return
	}
	response.OK(c, data)
}

func GetOrders(c *gin.Context) {
	var rowCount int64
	data := []model.Order{}

	paging := model.PagingInfo{}
	c.ShouldBindJSON(&paging)
	util.GetPaging(&paging)

	filter := "1=1"
	if !util.IsEmpty(paging.Filter) {
		filter = fmt.Sprintf("order_code LIKE '%%%[1]s%%' OR CAST(products AS char) LIKE '%%%[1]s%%' OR CAST(total_price AS char) LIKE '%%%[1]s%%'", paging.Filter)
	}

	res := util.DB.Model(&model.Order{}).Where(filter).Count(&rowCount).
		Order(util.StringJoin(paging.SortField, paging.SortOrder)).Limit(paging.Limit).Offset(paging.Offset).
		Find(&data)
	if res.Error != nil {
		if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
			response.AppError(c, res.Error.Error())
			return
		}
	}
	response.Paging(c, data, rowCount)
}

func FindOrderById(c *gin.Context) {
	var data = model.Order{}
	c.ShouldBindJSON(&data)

	if res := util.DB.First(&data, "id = ?", data.ID); res.Error != nil {
		response.Error(c, http.StatusNotFound, "Data not found!")
		return
	}
	response.OK(c, data)
}

func CreateOrder(c *gin.Context) {
	var req model.Order
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// uuidExists := true
	// for uuidExists {
	// 	req.OrderCode = util.GetNewUuid();
	// 	if res := util.DB.First(&req, "order_code = ?", req.OrderCode); res.Error != nil {
	// 		uuidExists = false
	// 	}
	// }

	var count int64 = 1
	for count > 0 {
		req.OrderCode = util.GetNewUuid()
		if res := util.DB.Model(&model.Order{}).Where("order_code = ?", req.OrderCode).Count(&count); res.Error != nil {
			response.AppError(c, res.Error.Error())
			return
		}
	}

	if res := util.DB.Create(&req); res.Error != nil {
		response.AppError(c, res.Error.Error())
		return
	}
	response.Void(c)
}

func UpdateOrder(c *gin.Context) {
	var req = model.Order{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var data = model.Order{}
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

func DeleteOrder(c *gin.Context) {
	var req = model.Order{}
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
