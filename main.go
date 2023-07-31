package main

import (
	"github.com/chanderah/menu-go/config"
	"github.com/gin-gonic/gin"
)
func main()  {
	config.ConnectDb();

	router:= gin.Default();
	router.Run("localhost:8080");
}