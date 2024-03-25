package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const tokenExpireTime = time.Hour * 72

var jwtSecret = []byte("明天会更好")

type JWTClaims struct {
	UserID    uint
	UserEmail string
	jwt.StandardClaims
}

// GenerateToken 签发用户 token
func GenerateToken(userID uint, userEmail string) (string, error) {
	claims := JWTClaims{
		UserID:    userID,
		UserEmail: userEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpireTime).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)

	return tokenString, err
}

// ParseToken 验证用户 token
func ParseToken(token string) (*JWTClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*JWTClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
