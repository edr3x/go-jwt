package routes

import (
	"github.com/gin-gonic/gin"

	"edr3x/go-jwt/controllers"
    "edr3x/go-jwt/middlewares"
)

func AuthRoute(r *gin.Engine){
    r.POST("/signup", controllers.SignUp)

    r.POST("/login", controllers.Login)

    r.GET("/validatedPath", middlewares.RequireAuth, controllers.Validate)
}
