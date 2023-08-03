package main

import (
	"net/http"

	"github.com/chanderah/menu-go/controller"
	"github.com/chanderah/menu-go/middleware"
	"github.com/chanderah/menu-go/util"
	"github.com/gin-gonic/gin"
)

func main() {
	// util.GetConnectionMySql()
	util.GetConnectionPostgres()

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
	apiRouter.Use(middleware.GinLoggingMiddleware)
	{
		postRouter := apiRouter.Group("/post")
		postRouter.POST("/create", controller.CreatePost)
	}
	{
		userRouter := apiRouter.Group("/user")
		userRouter.POST("/", controller.GetUsers)
		userRouter.POST("/register", controller.RegisterUser)
		userRouter.POST("/login", controller.LoginUser)

		userRouter.POST("/findById", controller.FindUserById)
		userRouter.POST("/findByUsername", controller.FindUserByUsername)
		userRouter.POST("/update", controller.UpdateUser)
		userRouter.POST("/delete", controller.DeleteUser)
	}
	router.Run(":" + port)
}
