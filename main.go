package main

// CompileDaemon -command="./go-auth-app"

import (
	"go-auth-app/controllers"
	"go-auth-app/initializers"
	"go-auth-app/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/sign-up", controllers.SignUp)
	r.POST("/sign-in", controllers.SignIn)
	r.GET("/validate", middleware.Auth, controllers.Validate)

	r.Run()
}
