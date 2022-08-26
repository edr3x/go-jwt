package controllers

import "github.com/gin-gonic/gin"

func Test(c *gin.Context) {
    c.IndentedJSON(200,gin.H{
        "message": "Success",
    })
}
