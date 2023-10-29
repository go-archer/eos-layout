package jwt

import (
	"fmt"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const jwtKey = "3ntMcJ"

// CreateToken 创建Token
func CreateToken(data map[string]string, keys ...string) string {
	key := jwtKey
	if len(keys) > 0 {
		key = keys[0]
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["iat"] = strconv.FormatInt(time.Now().Unix(), 10)
	claims["exp"] = strconv.FormatInt(time.Now().Add(time.Hour*2).Unix(), 10) // 2H
	for index, value := range data {
		claims[index] = value
	}
	token.Claims = claims
	tokenString, _ := token.SignedString([]byte(key))
	return tokenString
}

// ParseToken 解析token
func ParseToken(tokenStr string, keys ...string) (map[string]string, bool) {
	key := jwtKey
	if len(keys) > 0 {
		key = keys[0]
	}
	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		data := make(map[string]string)
		for index, val := range claims {
			data[index] = fmt.Sprintf("%v", val)
		}
		return data, true
	} else {
		return nil, false
	}
}
