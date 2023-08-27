package main

import (
	"github.com/chanderah/menu-go/controller"
	"github.com/chanderah/menu-go/middleware"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
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
		productRouter.POST("/findById", controller.FindProductById)
		productRouter.POST("/findByCategory", controller.FindProductByCategory)
		productRouter.POST("/findActiveByCategory", controller.FindActiveProductByCategory)
		productRouter.POST("/create", controller.CreateProduct)
		productRouter.POST("/update", controller.UpdateProduct)
		productRouter.POST("/delete", controller.DeleteProduct)
	}
	return router
}
