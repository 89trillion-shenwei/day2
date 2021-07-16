package router

import (
	"day2/internal/ctrl"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/Counter", ctrl.CountApi)
	return router
}
