package controllers

import (
	"mytokulist/database"
	"mytokulist/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	db := database.DbConnection
	var body struct {
		Username string
		Email    string
		Password string
	}

	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	user := models.Users{Username: body.Username, Email: body.Email, Password: string(hash)}
	sql := "INSERT INTO users (username, email, password, role) VALUES ($1, $2, $3, $4)"
	err = db.QueryRow(sql, user.Username, user.Email, user.Password, "User").Err()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Register",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success to Create User",
	})

}

func Login(c *gin.Context) {
	db := database.DbConnection
	var body struct {
		Username string
		Password string
	}

	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	var user = models.Users{}

	//scan into db
	sql := "SELECT id, username, password, role FROM users WHERE username = $1"
	err := db.QueryRow(sql, body.Username).Scan(&user.ID, &user.Username, &user.Password, &user.Role)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Username or Password is invalid",
		})
		panic(err)
	}

	//compare password with db password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Username or Password is invalid",
		})
		panic(err)
	}

	//make new jwt claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 12).Unix(),
	})

	//make new token
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Failed to create token",
		})
		panic(err)
	}

	//make cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*12, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})

}

func Logout(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}

	c.SetCookie("Authorization", tokenString, -1, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "You're logged out",
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message":   "Logged in as",
		"user_data": user,
	})
}

func NewPassword(c *gin.Context) {
	db := database.DbConnection
	var body struct {
		Password string
	}

	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	user, _ := c.Get("user")
	ok := bcrypt.CompareHashAndPassword([]byte(user.(models.Users).Password), []byte(body.Password))
	if ok == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password is same as old password.",
		})
		c.Abort()
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	sql := "UPDATE Users SET password = $1, updated_at = $ 2 WHERE id = $3"

	err = db.QueryRow(sql, hash, time.Now().Format(time.RFC3339), user.(models.Users).ID).Err()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update password",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success to update password",
	})
}
