package router

import (
	"log"
	"net/http"

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

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //請求頭部
		if origin != "" {
			//接收客戶端傳送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//伺服器支援的所有跨域請求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允許跨域設定可以返回其他子段，可以自定義欄位
			// c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
			// 允許瀏覽器（客戶端）可以解析的頭部 （重要）
			// c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//設定快取時間
			// c.Header("Access-Control-Max-Age", "172800")
			//允許客戶端傳遞校驗資訊比如 cookie (重要)
			// c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允許型別校驗
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery(), gin.Logger(), errorHandlingMiddleWare(logger.Error), Cors())

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
