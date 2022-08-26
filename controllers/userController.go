package controllers

import (
	"edr3x/go-jwt/initializers"
	"edr3x/go-jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Test
func Test(c *gin.Context) {
    c.IndentedJSON(200,gin.H{
        "message": "Success",
    })
}

// Sign up Controller
func SignUp(c *gin.Context){
    // Get the email and password from request body
    var body struct{
        Email string
        Password string
    }
    
    if c.Bind(&body) != nil{
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Failed to read body",
        })        
        return
    }

    // Hash the password
    hashedPass, err := bcrypt.GenerateFromPassword([]byte(body.Password),10)

    if err != nil{
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Failed to hash password",
        })
        return
    }

    //Create the user
    user := models.User{Email: body.Email, Password: string(hashedPass)}
    result := initializers.DB.Create(&user)

    if result.Error != nil{
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Failed to create user",
        })
    }
    
    // Response
    c.JSON(http.StatusOK, gin.H{
        "message": "User Created Successfully",
    })
}

