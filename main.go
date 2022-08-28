package main

import (
	"edr3x/go-jwt/config"
	"edr3x/go-jwt/routes"

	"github.com/gin-gonic/gin"
)

func init(){ 
   config.LoadEnv()
   config.ConnectToDB()
   config.DbSync()
}

func main() {
	r := gin.Default()
    
    routes.CheckRoute(r)
	routes.AuthRoute(r)

	r.Run()

}
