package router

import (
	"log"

	"github.com/borischen0203/Go-todolist/handlers"
	"github.com/borischen0203/Go-todolist/logger"

	// _ "github.com/borischen0203/Go-todolist/docs"

	"github.com/gin-gonic/gin"
	// ginSwagger "github.com/swaggo/gin-swagger"
	// "github.com/swaggo/gin-swagger/swaggerFiles"
)

var Router *gin.Engine

func errorHandlingMiddleWare(log *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err == nil {
			return
		}
		log.Printf("unexpected error: %s\n", err.Error())
	}
}

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery(), gin.Logger(), errorHandlingMiddleWare(logger.Error))

	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/health", handlers.HealthHandler)
	router.GET("/version", handlers.VersionHandler)

	router.POST("/todo", handlers.CreateItem)
	router.POST("/todo/:id", handlers.UpdateItem)
	router.GET("/todo-completed", handlers.GetCompletedItems)
	router.GET("/todo-incomplete", handlers.GetIncompleteItems)
	router.DELETE("/todo/:id", handlers.DeleteItem)

	return router
}

func Setup() {
	Router = SetupRouter()
}
