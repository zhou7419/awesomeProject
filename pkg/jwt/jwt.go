package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type CustomClaims struct {
	UserID    int    `json:"user_id"`
	UserName  string `json:"user_name"`
	ExpiresAt int64  `json:"expire_time"`
	jwt.RegisteredClaims
}

func Get(userId int, userName string) (string, error) {
	key := []byte("")
	expireAt := time.Now().Add(24 * time.Hour).Unix()
	claims := CustomClaims{
		UserID:           userId,
		UserName:         userName,
		ExpiresAt:        expireAt,
		RegisteredClaims: jwt.RegisteredClaims{},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString(key)

	return token, err
}
