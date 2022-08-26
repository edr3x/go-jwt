package main

import (
	"github.com/gin-gonic/gin"

	"edr3x/go-jwt/controllers"
	"edr3x/go-jwt/initializers"
    "edr3x/go-jwt/middlewares"
)

func init(){
   initializers.LoadEnv() 
   initializers.ConnectToDB()
   initializers.DbSync()
}

func main() {
    r := gin.Default()

    r.GET("/test", controllers.Test)

    r.POST("/signup", controllers.SignUp)

    r.POST("/login", controllers.Login)

    r.GET("/validatedPath", middlewares.RequireAuth, controllers.Validate)

    r.Run()

}
