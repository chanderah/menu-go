package main

import (
	"context"
	"fmt"
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
	// controller.SendMail("Http listen for http://go.chandrasa.fun started!", "aaa")
	// controller.SendMail("Http listen for http://go.chandrasa.fun started!", fmt.Sprintf("Your pid is: %d", os.Getpid()))
}

func getRouter() *gin.Engine {
	router := gin.New()
	router.Use(middleware.CorsMiddleware)
	// router.Use(middleware.SecurityMiddleware)

	/* MAIN API ROUTE */
	apiRouter := router.Group("/api")
	{
		fileRouter := apiRouter.Group("/file")
		fileRouter.POST("/findAll", controller.GetFiles)
		fileRouter.POST("/upload", controller.UploadFile)
		fileRouter.POST("/delete", controller.DeleteFile)
	}
	{
		userRouter := apiRouter.Group("/user")
		userRouter.POST("/findAll", controller.GetUsers)
		userRouter.POST("/findById", controller.FindUserById)
		userRouter.POST("/findByUsername", controller.FindUserByUsername)

		userRouter.POST("/register", controller.RegisterUser)
		userRouter.POST("/login", controller.LoginUser)
		userRouter.POST("/update", controller.UpdateUser)
		userRouter.POST("/delete", controller.DeleteUser)
	}
	{
		categoryRouter := apiRouter.Group("/category")
		categoryRouter.POST("/findAll", controller.GetCategories)
		categoryRouter.POST("/findById", controller.FindCategoryById)
		categoryRouter.POST("/create", controller.CreateCategory)
		categoryRouter.POST("/update", controller.UpdateCategory)
		categoryRouter.POST("/delete", controller.DeleteCategory)
	}
	{
		productRouter := apiRouter.Group("/product")
		productRouter.POST("/findAll", controller.GetProducts)
		productRouter.POST("/findActive", controller.GetActiveProducts)
		productRouter.POST("/findFeatured", controller.GetFeaturedProducts)
		productRouter.POST("/findById", controller.FindProductById)
		productRouter.POST("/findByCategory", controller.FindProductByCategory)
		productRouter.POST("/findActiveByCategory", controller.FindActiveProductByCategory)
		productRouter.POST("/findActiveByCategoryParam", controller.FindActiveProductByCategoryParam)
		productRouter.POST("/create", controller.CreateProduct)
		productRouter.POST("/update", controller.UpdateProduct)
		productRouter.POST("/delete", controller.DeleteProduct)
	}
	return router
}

func serve() {
	defer func() {
		if err := recover(); err != nil { //1
			fmt.Println("Gentle recovery from panic: %w", err)
		}
	}()

	util.GetConnectionMySql()
	// util.GetConnectionPostgres()

	port := "3000"
	router := getRouter()
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	{
		appRouter := router.Group("/app")
		// appRouter.POST("/mail", controller.CallSendMail)
		appRouter.GET("/info", func(c *gin.Context) {
			response.OK(c, map[string]interface{}{
				"pid": os.Getpid(),
			})
		})
		appRouter.GET("/kill/:pid", func(c *gin.Context) {
			pid, _ := strconv.Atoi(c.Param("pid"))
			if pid != os.Getpid() {
				response.Error(c, 400, "invalid!")
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
