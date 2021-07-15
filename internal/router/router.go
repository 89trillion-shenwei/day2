package router

import (
	handler2 "day2/internal/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router :=gin.Default()

	router.GET("/Counter", handler2.CountApi)
	return router
}