package routes

import (
	handlers "Hackmate/Handlers"
	middlewares "Hackmate/Middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	openRoutes := router.Group("api/o1")
	userAPI := openRoutes.Group("/user")
	userAPI.POST("/createProfile", handlers.CreateUserProfile)
	userAPI.POST("/login", handlers.Login)

	protectedRoutes := userAPI.Group("private")
	protectedRoutes.Use(middlewares.Authenticate())
	protectedRoutes.GET("/getdetails", handlers.GetUserDetails)
}
