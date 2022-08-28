package middlewares

import (
	"edr3x/go-jwt/config"
	"edr3x/go-jwt/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context){
    // Get the cookie from the request
    tokenString, err := c.Cookie("Authorization")

    if err != nil{
        c.AbortWithStatus(http.StatusUnauthorized)
        c.IndentedJSON(http.StatusUnauthorized,gin.H{
            "error":"Unauthenticated",
        })
    }
    
    // Decode  and validate the cookie
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	     	return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	   }

    	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
    	return []byte(os.Getenv("SECRET")), nil
    })

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        // Check the expire time
        if float64(time.Now().Unix()) > claims["exp"].(float64){
            c.AbortWithStatus(http.StatusUnauthorized)
        }

        // Find the user from token
        var user models.User
        config.DB.First(&user, claims["sub"])

        if user.ID == 0{
            c.AbortWithStatus(http.StatusUnauthorized)
            c.IndentedJSON(http.StatusUnauthorized,gin.H{
                 "error":"Unauthenticated",
        })

        }

        // Attach to request
        c.Set("user", user)

        // Continue
        c.Next()

    } else {
    	c.AbortWithStatus(http.StatusUnauthorized)
        c.IndentedJSON(http.StatusUnauthorized,gin.H{
            "error":"Unauthenticated",
        })

    }
}

