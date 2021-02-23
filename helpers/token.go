package helpers

import (
	"time"

	"github.com/form3tech-oss/jwt-go"
)

// ValidateToken dummy
func ValidateToken(token *jwt.Token, id string) string {
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	exp := int64(claims["exp"].(float64))
	if userID != id {
		return "Invalid token"
	}
	if exp < time.Now().Local().Unix() {
		return "Token has expired"
	}
	return "ok"
}
