package controller

import (
	"log"
	"net/http"

	"github.com/chanderah/menu-go/config"
	"github.com/chanderah/menu-go/model"
	"github.com/gin-gonic/gin"
)

func All(c *gin.Context) {
	post := []model.Post{}
	if err := config.DB.Select(&post, "SELECT * FROM tutorial.posts"); err != nil {
		log.Println(err);
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "not found"})
	}
	// c.JSON(http.StatusOK, post)
	c.JSON(http.StatusOK, gin.H{"data": post})
}

func FindById(c *gin.Context) {
	post := []model.Post{}
	if err := config.DB.Select(&post, "SELECT * FROM tutorial.posts"); err != nil {
		log.Println(err);
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "not found"})
	}
	c.JSON(http.StatusOK, gin.H{"data": post})
}

func Create(c *gin.Context) {
	post := []model.Post{}
	if err := config.DB.Select(&post, "SELECT * FROM tutorial.posts"); err != nil {
		log.Println(err);
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "not found"})
	}
	c.JSON(http.StatusOK, gin.H{"data": post})
}

func Update(c *gin.Context) {
	post := []model.Post{}
	if err := config.DB.Select(&post, "SELECT * FROM tutorial.posts"); err != nil {
		log.Println(err);
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "not found"})
	}
	c.JSON(http.StatusOK, gin.H{"data": post})
}

func Delete(c *gin.Context) {
	post := []model.Post{}
	if err := config.DB.Select(&post, "SELECT * FROM tutorial.posts"); err != nil {
		log.Println(err);
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "not found"})
	}
	c.JSON(http.StatusOK, gin.H{"data": post})
}