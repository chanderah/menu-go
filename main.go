package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/chanderah/menu-go/response"
	"github.com/chanderah/menu-go/util"
	"github.com/gin-gonic/gin"
)

func main() {
	serve()
	// controller.SendMail("Http listen for http://go.chandrasa.fun started!", "aaa")
	// controller.SendMail("Http listen for http://go.chandrasa.fun started!", fmt.Sprintf("Your pid is: %d", os.Getpid()))
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
	router := GetRouter()
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
