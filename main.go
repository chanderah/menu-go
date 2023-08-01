package main

import (
	"net/http"

	"github.com/chanderah/menu-go/controller"
	"github.com/chanderah/menu-go/middleware"
	"github.com/chanderah/menu-go/util"
	"github.com/gin-gonic/gin"
)

func main() {
	util.GetConnectionMySql()

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
	/* MAIN ROUTE */

	postRouter := apiRouter.Group("/post")
	postRouter.POST("/create", controller.CreatePost)

	userRouter := apiRouter.Group("/user")
	userRouter.POST("/", controller.GetUsers)
	userRouter.POST("/register", controller.RegisterUser)
	userRouter.POST("/login", controller.LoginUser)

	userRouter.POST("/findById", controller.FindUser)
	userRouter.POST("/update", controller.UpdateUser)
	userRouter.POST("/delete", controller.DeleteUser)

	// userRouter := apiRouter.Group("/user")
	// userRouter.POST("/", controller.GetUsers)
	// userRouter.POST("/findById", controller.FindById)
	// userRouter.POST("/create", controller.CreateUser)
	// userRouter.POST("/update", controller.UpdateUser)
	// userRouter.POST("/delete", controller.DeleteUser)

	router.Run(":" + port)
}
