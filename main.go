package main

import (
	"net/http"

	"github.com/chanderah/menu-go/config"
	"github.com/chanderah/menu-go/controller"
	"github.com/gin-gonic/gin"
)

func main()  {
	config.ConnectDb();

	server:= ":3001";
	router:= gin.Default();

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"message": "Welcome!",
		})
		});

	apiRouter:= router.Group("/api")

	postRouter:= apiRouter.Group("/post");
	postRouter.POST("/create", controller.Create)

	router.Run(server);
}