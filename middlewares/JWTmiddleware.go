package middlewares

import (
	"fajar/clean/config"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMidlleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(config.SECRET_JWT), ///pencocokan data valid atau tidak
	})
}
func CreateToken(userId int, role string) (string, error) { //membuat token jwt
	claims := jwt.MapClaims{} //membuat sebuah map di jwt yang didalamnya berisi informasi payload data user
	claims["authorized"] = true
	claims["userId"] = userId
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //token expired setelah 1 jam
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.SECRET_JWT))
}

func ExtractTokenUserId(e echo.Context) int { // mengextrak informasi yang ada di token
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return int(userId)
	}
	return 0

}
func ExtractTokenUserRole(e echo.Context) string { // mengextrak informasi yang ada di token
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		role := claims["role"].(string)
		return role
	}
	return ""

}
