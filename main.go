package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"

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
	util.GetConnectionMySql()
	// util.GetConnectionPostgres()

	port := "3001"
	router := generateRoute()
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	{
		appRouter := router.Group("/app")
		appRouter.GET("/info", func(c *gin.Context) {
			response.OK(c, map[string]interface{}{
				"pid": os.Getpid(),
			})
		})
		appRouter.GET("/kill/:pid", func(c *gin.Context) {
			pid, _ := strconv.Atoi(c.Param("pid"))
			if pid != os.Getpid() {
				response.Error(c, 400, "Invalid!")
				return
			}
			log.Println("Shutting down...")
			if err := srv.Shutdown(context.Background()); err != nil {
				log.Println("Stopping server failed.\n", err)
			}
		})
	}

	router.GET("/", func(c *gin.Context) {
		response.OK(c, "Welcome!")
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
		userRouter.POST("/findById", controller.FindUserById)
		userRouter.POST("/findByCategory", controller.FindProductByCategory)
		userRouter.POST("/findByUsername", controller.FindUserByUsername)

		userRouter.POST("/register", controller.RegisterUser)
		userRouter.POST("/login", controller.LoginUser)

		userRouter.POST("/update", controller.UpdateUser)
		userRouter.POST("/delete", controller.DeleteUser)
	}
	{
		productRouter := apiRouter.Group("/product")
		productRouter.POST("/", controller.GetProducts)
		productRouter.POST("/findById", controller.FindProductById)
		productRouter.POST("/findByCategory", controller.FindProductByCategory)

		productRouter.POST("/create", controller.CreateProduct)
		productRouter.POST("/update", controller.UpdateProduct)
		productRouter.POST("/delete", controller.DeleteProduct)
	}
	return router
}
