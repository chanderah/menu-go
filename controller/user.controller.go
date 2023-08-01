package controller

import (
	"net/http"

	"github.com/chanderah/menu-go/config"
	"github.com/chanderah/menu-go/model"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var data model.User
	if	err:= c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": err.Error(),
		})
		return;
	}

	config.Db.Create(&data);
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"message": "success",
		"data": data,
	})
}