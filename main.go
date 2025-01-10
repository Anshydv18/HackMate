package main

import (
	base "Hackmate/Base"
	middlewares "Hackmate/Middlewares"
	routes "Hackmate/Routes"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middlewares.SetContext())
	routes.UserRoutes(router)
	start := time.Now()
	base.Initiate()

	fmt.Println(time.Now().Sub(start))
	router.Run(":3000")
}
