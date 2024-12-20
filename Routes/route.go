package routes

import (
	handlers "NotesBuddy/Handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	openRoutes := router.Group("api/o1")

	userRoute := openRoutes.Group("/user")
	userRoute.POST("/createProfile", handlers.CreateUserProfile)

}
