package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/victorguidi/initializers"
	"github.com/victorguidi/models"
	"net/http"
	"os"
	"time"
)

func ValidateJwt(c *gin.Context) {
	// get the cookie off the request
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		// log.Println("yaaay")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	//validate Jwt
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		//Check the expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//find the user with the token sub
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// attach to req

		c.Set("user", claims["sub"])
		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
