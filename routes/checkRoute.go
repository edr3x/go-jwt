package routes

import (
	"edr3x/go-jwt/controllers"

	"github.com/gin-gonic/gin"
)

func CheckRoute(r *gin.Engine){
    r.GET("/", controllers.RouteCheck)
}
