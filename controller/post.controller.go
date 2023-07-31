package controller

import (
	"net/http"

	"github.com/chanderah/menu-go/config"
	"github.com/chanderah/menu-go/model"
	"github.com/gin-gonic/gin"
)

// type CPost struct {
// 	Title   string `json:"title" binding:"required"`
// 	Content string `json:"content" binding:"required"`
// }

func Create(c *gin.Context) {
	// var data CPost
	var data model.Post
	err:= c.ShouldBindJSON(&data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": err.Error(),
		})
		return;
	}

	post := model.Post{Title: data.Title, Content: data.Content};
	config.DB.Create(&post);

	c.JSON(http.StatusOK, gin.H{ "data": post })
	// response.Ok(post)
}