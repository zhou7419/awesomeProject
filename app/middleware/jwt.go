package middleware

import (
	pkgJwt "awesomeProject/pkg/jwt"
	"awesomeProject/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")
		token, err := jwt.ParseWithClaims(tokenStr, &pkgJwt.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			// 这个函数返回用于验证签名的密钥
			return []byte(""), nil
		})

		if err != nil {
			response.Unauthorized(c, "token验证失败")
			return
		}

		if !token.Valid {
			response.Unauthorized(c, "token验证失败")
			return
		}

		claims := token.Claims.(*pkgJwt.CustomClaims)
		c.Set("claims", claims)

		c.Next()
	}
}
