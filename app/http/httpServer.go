package http

import (
	router2 "day2/internal/router"
)

func Start() {
	router := router2.InitRouter()
	router.Run(":8000")

}
