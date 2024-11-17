package token

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

var (
	ErrSignString          = errors.New("jwt sign string err")
	ErrParseJwtToken       = errors.New("parse jwt token failed")
	ErrInvalidToken        = errors.New("invalid token")
	InvalidUserID    int32 = -1
)

type JWTer struct{}

func NewJWTer() *JWTer {
	return &JWTer{}
}

//使用jwt来生成token和验证token

func (j *JWTer) GenerateJwtToken(userID int32, jwtSecret string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = userID

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", ErrSignString
	}
	return tokenString, nil
}

func (j *JWTer) VerifyJwtToken(tokenString, jwtSecret string) (int32, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return InvalidUserID, ErrParseJwtToken
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return InvalidUserID, ErrInvalidToken

	}
	// 修改：断言为 float64 再转为 int32
	userIDFloat, ok := claims["userID"].(float64)
	if !ok {
		return InvalidUserID, ErrInvalidToken
	}

	userID := int32(userIDFloat) // 转换为 int32
	return userID, nil
}
