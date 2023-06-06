package helpers

import "github.com/golang-jwt/jwt/v5"

var secretKey = "rahasia123"

func GenerateToken(id uint, username string) string {
	claims := jwt.MapClaims{
		"id":       id,
		"username": username,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := parseToken.SignedString([]byte(secretKey))

	return signedToken
}
