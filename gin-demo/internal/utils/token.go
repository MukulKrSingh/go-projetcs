package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secret = "secret_code_for_notes_api"

func GenerateToken(email string, id int) (string, error) {
	fmt.Println(email)
	fmt.Println(id)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"id":    id,
			"nbf":   time.Date(2024, time.August, 8, 1, 1, 1, 1, time.UTC).Unix(),
		})
	tokenStr, err := token.SignedString([]byte(secret))
	fmt.Println(tokenStr)
	return tokenStr, err
}
