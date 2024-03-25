package main

import (
	"backWeb/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type MyCustomClaims struct {
	jwt.RegisteredClaims
}

func GenerateToken(user models.User) string {

	claims := MyCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
			ID:        strconv.Itoa(int(user.ID)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, _ := token.SignedString([]byte("TheMostSecureKeyInTheWorld")) //Esto no deberia estar aqui xd

	return tokenSigned
}

func ValidateToken(c *gin.Context) {
	cookie := c.GetHeader("Cookie")
	token, _ := jwt.Parse(cookie[6:], func(token *jwt.Token) (interface{}, error) {
		return []byte("TheMostSecureKeyInTheWorld"), nil //Esto si que menos deberia estar aqui, pero jaja xd
	})
	if token.Valid {
		fmt.Println("Valid token")
		c.Next()
	} else {
		fmt.Println("Invalid token")
		c.AbortWithStatus(401)

	}

}

func Logger(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		c.BindJSON(&user)

		result := db.First(&user, "username = ? AND password=?", user.Username, user.Password)
		if result.RowsAffected > 0 {
			fmt.Println("User found")
			tokenSigned := GenerateToken(user)
			c.SetCookie("Token", tokenSigned, 3600, "/", "", false, false)
			fmt.Println("tokenSigned: ", tokenSigned)
			c.Next()
		} else {
			fmt.Println("User not found")
			c.AbortWithStatus(401)

		}

	}

}
