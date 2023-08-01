package controller

import (
	"net/http"
	"time"

	"github.com/chanderah/menu-go/config"
	"github.com/chanderah/menu-go/model"
	"github.com/gin-gonic/gin"
)

type UserBasic struct {
	ID        uint      `json:"id"`
	Role      string    `json:"role"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func RegisterUser(c *gin.Context) {
	var data model.User
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if res := config.DB.Create(&data); res.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": res.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "success",
	})
}

func LoginUser(c *gin.Context) {
	var input, data model.User

	c.ShouldBindJSON(&input)
	if res := config.DB.First(&data, "username = ?", input.Username); res.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found!",
		})
		return
	}

	if input.Password != data.Password {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Bad Credentials!",
		})
		return
	}

	if res := config.DB.Model(&data).Updates(input); res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": res.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "success",
	})
}

func GetUsers(c *gin.Context) {
	var data []UserBasic
	if res := config.DB.Find(&[]model.User{}).Scan(&data); res.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": res.Error,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "success",
		"data":    data,
	})
}

func FindUser(c *gin.Context) {
	var data model.User

	c.ShouldBindJSON(&data)
	// if res:=config.DB.Raw("select * from users where id = ?", 3).Scan(&data)
	if res := config.DB.First(&data, "id = ?", data.ID); res.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "success",
		"data":    data,
	})
}

func UpdateUser(c *gin.Context) {
	var input, data model.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if res := config.DB.First(&data, "id = ?", input.ID); res.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found!",
		})
		return
	}

	if res := config.DB.Model(&data).Updates(input); res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": res.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "success",
		"data":    data,
	})
}

func DeleteUser(c *gin.Context) {
	var data model.User

	c.ShouldBindJSON(&data)
	if res := config.DB.First(&data, "id = ?", data.ID); res.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found!",
		})
		return
	}

	config.DB.Delete(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "success",
	})
}
