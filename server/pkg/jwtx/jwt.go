package jwtx

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

func ParseAndGetUserID(tokenString string, secretKey []byte) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法是否是期望的算法，这里假设使用HS256算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method.Alg())
		}
		return secretKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["Id "].(string)
		if ok {
			return userID, nil
		}
		return "", fmt.Errorf("user_id claim not found or not in string format")
	}
	return "", fmt.Errorf("invalid JWT token")
}
