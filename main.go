package main

import (
	"context"
	"log"
	"net/http"

	"github.com/chanderah/menu-go/controller"
	"github.com/chanderah/menu-go/middleware"
	"github.com/chanderah/menu-go/response"
	"github.com/chanderah/menu-go/util"
	"github.com/gin-gonic/gin"
)

func main() {
	serve()
}

func serve() {
	// util.GetConnectionMySql()
	util.GetConnectionPostgres()

	port := "3001"
	router := generateRoute()
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	router.GET("/", func(c *gin.Context) {
		response.OK(c, "Welcome!")
	})
	router.GET("/app/kill", func(c *gin.Context) {
		log.Println("Shutting down...")
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Println("Server is already closed.\n", err)
		}
	})

	if err := srv.ListenAndServe(); err != nil {
		log.Printf("listen: %s\n", err)
	}
}

func generateRoute() *gin.Engine {
	router := gin.New()
	/* MAIN API ROUTE */
	apiRouter := router.Group("/api")
	apiRouter.Use(middleware.GinLoggingMiddleware)
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
	{
		productRouter := apiRouter.Group("/product")
		productRouter.POST("/", controller.GetUsers)
		productRouter.POST("/register", controller.RegisterUser)
		productRouter.POST("/login", controller.LoginUser)

		productRouter.POST("/findById", controller.FindUserById)
		productRouter.POST("/findByUsername", controller.FindUserByUsername)
		productRouter.POST("/update", controller.UpdateUser)
		productRouter.POST("/delete", controller.DeleteUser)
	}
	return router
}
