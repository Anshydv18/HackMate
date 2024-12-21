package main

import (
	middlewares "NotesBuddy/Middlewares"
	routes "NotesBuddy/Routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middlewares.SetContext())
	routes.UserRoutes(router)

	router.Run(":8000")
}
