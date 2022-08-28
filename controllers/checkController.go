package controllers

import "github.com/gin-gonic/gin"

// Test
func RouteCheck(c *gin.Context) {
    c.IndentedJSON(200,gin.H{
        "message": "Up and Runnning",
    })
}


