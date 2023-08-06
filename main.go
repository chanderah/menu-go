package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// serve()
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":3001"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hellooo, World!")
}

func serve() {
	// util.GetConnectionMySql()
	// // util.GetConnectionPostgres()

	// port := "3001"
	// router := route()

	// router.GET("/", func(c *gin.Context) {
	// 	response.OK(c, "Welcome!")
	// })

	// router.GET("/hi", func(c *gin.Context) {
	// 	response.OK(c, "Hi!")
	// })

	// router.Run(":" + port)

	// if err := srv.ListenAndServe(); err != nil {
	// 	log.Printf("listen: %s\n", err)
	// }
}

func route() *gin.Engine {
	router := gin.New()
	/* MAIN API ROUTE */
	// apiRouter := router.Group("/api")
	// apiRouter.Use(middleware.GinLoggingMiddleware)
	// {
	// 	userRouter := apiRouter.Group("/user")
	// 	userRouter.POST("/", controller.GetUsers)
	// 	userRouter.POST("/register", controller.RegisterUser)
	// 	userRouter.POST("/login", controller.LoginUser)

	// 	userRouter.POST("/findById", controller.FindUserById)
	// 	userRouter.POST("/findByUsername", controller.FindUserByUsername)
	// 	userRouter.POST("/update", controller.UpdateUser)
	// 	userRouter.POST("/delete", controller.DeleteUser)
	// }
	// {
	// 	productRouter := apiRouter.Group("/product")
	// 	productRouter.POST("/", controller.GetUsers)
	// 	productRouter.POST("/register", controller.RegisterUser)
	// 	productRouter.POST("/login", controller.LoginUser)

	// 	productRouter.POST("/findById", controller.FindUserById)
	// 	productRouter.POST("/findByUsername", controller.FindUserByUsername)
	// 	productRouter.POST("/update", controller.UpdateUser)
	// 	productRouter.POST("/delete", controller.DeleteUser)
	// }
	return router
}
