package main

import (
	middlewares "Hackmate/Middlewares"
	routes "Hackmate/Routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middlewares.SetContext())
	routes.UserRoutes(router)

	router.Run(":3000")
}
