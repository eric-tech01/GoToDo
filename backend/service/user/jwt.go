package userSrv

import (
	"errors"
	"fmt"
	userM "netix-todo-app/backend/model/user"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = "netix-todo-app/backend"

func GenerateToken(userId int64) (string, error) {
	// 创建一个新的 Token 对象，并设置有效期为1小时
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": userId,
		"exp":    time.Now().Add(time.Duration(12) * time.Hour).Unix(),
	})
	// 使用 SecretKey 对 Token 进行签名
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// parse token
func ParseToken(tokenString string) (jwt.MapClaims, error) {
	var err error
	// 解析 Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// 返回密钥
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	// 验证解析结果
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("token is invalid")
	}
}

func VerifyToken(tokenString string) (*userM.User, string, error) {
	var err error
	// 解析 Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// 返回密钥
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, "", err
	}

	// 验证解析结果
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp := claims["exp"].(float64)
		expTime := time.Unix(int64(exp), 0)
		userId := int64(claims["userid"].(float64))

		var newToken string
		if expTime.Before(time.Now().Add(1 * time.Hour)) {
			//查询db获取用户信息
			newToken, err = GenerateToken(userId)
		}
		// if expTime.
		return &userM.User{
			Id: userId,
		}, newToken, err
	} else {
		return nil, "", errors.New("token is invalid")
	}
}
