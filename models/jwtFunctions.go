package models

import (
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

func GenerateToken(user User) string {

	claims := MyCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			ID:        strconv.Itoa(int(user.ID)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, _ := token.SignedString([]byte("TheMostSecureKeyInTheWorld")) //Esto no deberia estar aqui xd

	return tokenSigned
}

func ValidateToken(c *gin.Context) {
	cookie := c.GetHeader("Cookie")
	token, error := jwt.Parse(cookie[6:], func(token *jwt.Token) (interface{}, error) {
		return []byte("TheMostSecureKeyInTheWorld"), nil //Esto si que menos deberia estar aqui, pero jaja xd
	})
	if error != nil {
		fmt.Println("Error parsing token")
		c.AbortWithStatus(401)
	}
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
		var user User
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
func GetIdJWT(tokenString string) string {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Esto hace magia
		return []byte("TheMostSecureKeyInTheWorld"), nil
	})

	userId := token.Claims.(jwt.MapClaims)["jti"].(string)
	return userId
}
