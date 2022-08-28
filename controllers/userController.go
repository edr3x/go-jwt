package controllers

import (
	"edr3x/go-jwt/config"
	"edr3x/go-jwt/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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
    result := config.DB.Create(&user)

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

// Login 
func Login(c *gin.Context){
    // Get email and password from request body
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

    // Look up requested user in database
    var user models.User
    config.DB.First(&user, "email = ?", body.Email)

    if user.ID == 0{
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid Email or password",
        })
        return
    }

    // Compare password with password saved in database
    err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(body.Password))

    if err != nil{
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid Password",
        })
    }

    // Generate jwt token
    userToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "sub": user.ID,
        "exp": time.Now().Add(time.Hour * 24 *30).Unix(),
    })
    
    tokenString, err := userToken.SignedString([]byte(os.Getenv("SECRET")))

    if err != nil{
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Failed to create token",
        })
        return
    }

    // Set Cookie
    c.SetSameSite(http.SameSiteLaxMode)
    c.SetCookie("Authorization", tokenString, 3600 * 24 *30, "", "", false, true)
              //name           , value      , maxAge    ,path, domain,secure, httpOnly

    // Send jwt token back to user
    c.IndentedJSON(http.StatusOK,gin.H{})
}

func Validate(c *gin.Context){
    user, _ := c.Get("user")

    c.JSON(http.StatusOK, gin.H{
        "message": user,
    })
}
