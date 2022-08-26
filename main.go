package main

import (
	"github.com/gin-gonic/gin"

	"edr3x/go-jwt/controllers"
	"edr3x/go-jwt/initializers"
)

// This branch is for future to set up server faster

func init(){
   initializers.LoadEnv() 
   initializers.ConnectToDB()
   initializers.DbSync()
}

func main() {
    r := gin.Default()

    r.GET("/test", controllers.Test)

    r.Run()

}
