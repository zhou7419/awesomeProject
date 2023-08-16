package middleware

import (
	"awesomeProject/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenStr := c.Request.Header.Get("Authorization")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			// 在这里返回用于验证签名的密钥
			return []byte(""), nil
		})

		if err != nil {
			response.Unauthorized(c, "token验证失败")
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("claims", claims)
		} else {
			response.Unauthorized(c, "token验证失败")
			return
		}

		c.Next()
	}
}
