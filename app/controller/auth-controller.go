package controller

import (
	// "encoding/json"
	// "fmt"
	"net/http"
	"online-store/app/db"
	"online-store/app/model"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"

	// "online-store/utils"
)

func Register(c *gin.Context) {
	var userRequest model.UserRegisterRequest

	//validate
	if err := c.Bind(&userRequest); err != nil {
		if err.Error() == "Key: 'UserRegisterRequest.Username' Error:Field validation for 'Username' failed on the 'required' tag" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
			return
		}
		if err.Error() == "Key: 'UserRegisterRequest.Email' Error:Field validation for 'Email' failed on the 'required' tag" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
			return
		}
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password.",
		})
		return
	}

	user := model.User{
		Username: userRequest.Username,
		Email:    userRequest.Email,
		Password: string(hash),
	}

	tx := db.DBConn.Begin()

	if err := tx.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tx.Commit()

	userResponse := model.UserResponse{
		Username: user.Username,
		Email:    user.Email,
	}

	c.JSON(http.StatusOK, model.UserRegisterResponse{
		Message: "User has been created",
		Data:    userResponse,
	})

}

func Login(c *gin.Context) {
	var userRequest model.UserLoginRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var user model.User

	tx := db.DBConn.Begin()

	tx.Where("email = ?", userRequest.Email).First(&user)
	if err := tx.Where("email = ?", userRequest.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email or password"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"nessage": "Invalid email or password"})
		return
	}

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Minute * 1).Unix(), // satu menit
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create token"})
		return
	}

	// Respond
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	userResponse := model.UserResponse{
		Username: user.Username,
		Email:    user.Email,
	}

	c.JSON(http.StatusOK, model.UserLoginResponse{
		Data:  userResponse,
		Token: tokenString,
	})
}

// type TokenResponse struct {
// 	AccessToken string `json:"access_token"`
// }

// func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
// 	// Ambil refresh token dari header atau body request
// 	refreshToken := r.Header.Get("Authorization")
// 	if refreshToken == "" {
// 		http.Error(w, "Refresh token is required", http.StatusBadRequest)
// 		return
// 	}

// 	// Validasi refresh token
// 	claims, err := utils.ValidateToken(refreshToken)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Invalid refresh token: %v", err), http.StatusUnauthorized)
// 		return
// 	}

// 	// Jika valid, buat access token baru
// 	newAccessToken, err := utils.GenerateAccessToken(claims.UserID)
// 	if err != nil {
// 		http.Error(w, "Could not create new access token", http.StatusInternalServerError)
// 		return
// 	}

// 	// Return token baru dalam format JSON
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(TokenResponse{AccessToken: newAccessToken})
// }

