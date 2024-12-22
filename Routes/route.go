package routes

import (
	handlers "NotesBuddy/Handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	openRoutes := router.Group("api/o1")
	userAPI := openRoutes.Group("/user")
	userAPI.POST("/createProfile", handlers.CreateUserProfile)
	userAPI.POST("/login", handlers.Login)

}
