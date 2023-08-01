package controller

import (
	"net/http"

	"github.com/chanderah/menu-go/config"
	"github.com/chanderah/menu-go/model"
	"github.com/gin-gonic/gin"
)

// type User struct {
// 	model.User
// 	Username string `json:"username" binding:"required"`
// 	Password string `json:"password" binding:"required"`
// }

func GetUsers(c *gin.Context) {
	// var data []User
	var data []model.User
	if res := config.DB.Find(&data); res.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": res.Error,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"message": "success",
		"data": data,
	})
}

func FindById(c *gin.Context) {
	var data model.User;

	c.ShouldBindJSON(&data);
	// if res:=config.DB.Raw("select * from users where id = ?", 3).Scan(&data)
	if res:= config.DB.First(&data, "id = ?", data.ID); res.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"message": "Data not found!",
		})
		return;
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"message": "success",
		"data": data,
	})
}

func CreateUser(c *gin.Context) {
	var data model.User
	if err:= c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": err.Error(),
		})
		return;
	}

	config.DB.Create(&data);
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"message": "success",
		"data": data,
	})
}

func UpdateUser(c *gin.Context) {
	var data model.User
	if err:= c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": err.Error(),
		})
		return;
	}

	config.DB.Create(&data);
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"message": "success",
		"data": data,
	})
}

func DeleteUser(c *gin.Context) {
	var data model.User
	if err:= c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": err.Error(),
		})
		return;
	}

	config.DB.Create(&data);
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"message": "success",
		"data": data,
	})
}