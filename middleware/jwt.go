package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/codycoding/goDuck/core"
	"github.com/codycoding/goDuck/global"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"time"
)

const TokenRedisKeyPrefix = "LoginToke:"

// 数据结构定义

//
// UserInfo
//  @Description: 用户登录信息
//
type UserInfo struct {
	AccountId int64  `json:"accountId"` // 用户ID
	UserID    int64  `json:"userId"`    // 用户ID(临时)
	UserName  string `json:"userName"`  // 用户名称
	NickName  string `json:"nickName"`  // 用户昵称
	RoleId    string `json:"roleId"`    // 角色ID
}

//
// CustomClaims
//  @Description: 自定义Token内容结构
//
type CustomClaims struct {
	UserInfo           UserInfo // 用户登录信息
	jwt.StandardClaims          // 标准载荷信息
}

var (
	TokenExpired     = errors.New("过期Token")
	TokenNotValidYet = errors.New("未生效Token")
	TokenMalformed   = errors.New("无效Token")
	TokenInvalid     = errors.New("非法Token")
)

// JwtAuth JWT鉴权中间件
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("SlpAuthorization")
		if token == "" {
			// 无token信息，返回203
			core.UnauthorizedWithMessage("未登录或非法访问", c)
			c.Abort()
			return
		}
		// token解析
		claims, err := ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				// token信息已过期，返回203
				core.UnauthorizedWithMessage("授权已过期", c)
				c.Abort()
				return
			}
			// 其他错误
			// 返回203
			core.UnauthorizedWithMessage(err.Error(), c)
			c.Abort()
			return
		}
		// 匹配Redis
		if CompareRedisToken(claims.UserInfo, token) {
			// 匹配成功
			c.Set("claims", claims)
			c.Next()
		} else {
			// 匹配失败
			// 无效Token，返回203
			core.UnauthorizedWithMessage(TokenInvalid.Error(), c)
			c.Abort()
			return
		}
	}
}

// ParseToken token解析方法，返回含有用户信息的数据结构
func ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.JWT.SigningKey), nil
	})
	// 错误类型判断
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				// 无效Token
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// 过期Token
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				// 未生效Token
				return nil, TokenNotValidYet
			} else {
				// 其他错误
				// 非法Token
				return nil, TokenInvalid
			}
		} else {
			return nil, err
		}
	}
	// token内容判断
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			// token解析成功
			return claims, nil
		}
		// 数据结构解析失败
		return nil, TokenInvalid
	} else {
		// token为空
		return nil, TokenInvalid
	}
}

// GetTokenRedisKey 获取token缓存Key
func GetTokenRedisKey(userId int64, userName string) string {
	return fmt.Sprintf("%s%d-%s", TokenRedisKeyPrefix, userId, userName)
}

// CompareRedisToken 当前Token与Redis记录Token匹配
func CompareRedisToken(userInfo UserInfo, tokenStr string) bool {
	redisKey := GetTokenRedisKey(userInfo.AccountId, userInfo.UserName)
	ctx := context.Background()
	redisToken, err := global.Redis.Get(ctx, redisKey).Result()
	if err == redis.Nil {
		// redis没有token记录
		return false
	} else if err != nil {
		// redis获取出错
		return false
	}
	return redisToken == tokenStr
}

// CreateToken 根据用户信息生成Token
func CreateToken(userInfo UserInfo) (string, error) {
	// jwt签名bytes
	var signingKey = []byte(global.Config.JWT.SigningKey)
	// token载体生成
	claims := CustomClaims{
		UserInfo: userInfo,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(global.Config.JWT.ExpiresTime) * time.Minute).Unix(), // 签名过期时间
			Issuer:    "slp",                                                                             // 签名发行者
			NotBefore: time.Now().Unix(),                                                                 // 签名生效时间 立即生效
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// token字符串生成并写入redis
	if tokenStr, sErr := token.SignedString(signingKey); sErr != nil {
		return "", sErr
	} else {
		ctx := context.Background()
		redisKey := GetTokenRedisKey(userInfo.AccountId, userInfo.UserName)
		global.Redis.Set(ctx, redisKey, tokenStr, time.Duration(global.Config.JWT.ExpiresTime)*time.Minute)
		return tokenStr, nil
	}
}
