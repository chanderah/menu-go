package main

import (
	"net/http"

	"github.com/chanderah/menu-go/config"
	"github.com/chanderah/menu-go/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDb()

	port := "3001"
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "Welcome!",
		})
	})

	/* MAIN ROUTE */
	apiRouter := router.Group("/api")
	/* MAIN ROUTE */

	postRouter := apiRouter.Group("/post")
	postRouter.POST("/create", controller.CreatePost)

	userRouter := apiRouter.Group("/user")
	userRouter.POST("/", controller.GetUsers)
	userRouter.POST("/findById", controller.FindById)
	userRouter.POST("/create", controller.CreateUser)
	userRouter.POST("/update", controller.UpdateUser)
	userRouter.POST("/delete", controller.DeleteUser)

	router.Run(":" + port)
}
