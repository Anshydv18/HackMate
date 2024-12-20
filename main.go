package main

import (
	routes "NotesBuddy/Routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.UserRoutes(router)

	router.Run(":8000")
}
