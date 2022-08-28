package main

import (
	"github.com/gin-gonic/gin"

	"edr3x/go-jwt/controllers"
	"edr3x/go-jwt/config"
    "edr3x/go-jwt/middlewares"
)

func init(){
   config.LoadEnv()
   config.ConnectToDB()
   config.DbSync()
}

func main() {
    r := gin.Default()

    r.GET("/test", controllers.Test)

    r.POST("/signup", controllers.SignUp)

    r.POST("/login", controllers.Login)

    r.GET("/validatedPath", middlewares.RequireAuth, controllers.Validate)

    r.Run()

}
