package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type CustomClaims struct {
	UserID   int64
	UserName string
	jwt.StandardClaims
}

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token已过期")
	TokenNotValidYet = errors.New("Token未生效")
	TokenMalformed   = errors.New("Token无效")
	TokenInvalid     = errors.New("非法Token:")
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Header携带Token验证
		//if token := c.Request.Header.Get("Authorization"); token == "" {
		//	// 无Token访问
		//	core.FailWithMessage(http.StatusUnauthorized, "登录失效", c)
		//	global.SlpLog.Error("未提供Token信息。", zap.String("ip", c.ClientIP()))
		//	c.Abort()
		//} else {
		//	// TODO next
		//}
		c.Next()
	}
}

//
// parseToken
//  @Description: 解析Token
//  @receiver j
//  @param tokenString
//  @return *CustomClaims
//  @return error
//
func (j *JWT) parseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid

	}
}
