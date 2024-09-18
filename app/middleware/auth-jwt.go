package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	// Get the cookie off the request
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		//error response
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Decode/validate it
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Chec k the expiry date
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// // Find the user with token Subject
		// var user model.User
		// if err := db.DBConn.First(&user, claims["sub"]).Error; err != nil {
		// 	c.AbortWithStatus(http.StatusUnauthorized)
		// }

		// // Attach the request
		// c.Set("user",user)

		//Continue
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		c.AbortWithStatus(http.StatusUnauthorized)

	}
}
