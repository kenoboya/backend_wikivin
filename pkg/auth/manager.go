package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrSecretKeyIsEmpty = errors.New("secret key is empty")
	ErrUnauthorized = errors.New("token is invalid")
)

type TokenManager interface {
	NewJWT(userID int, role string, ttl time.Duration) (string, error)
	VerifyToken(token string) (int, error)
	RefreshToken(token string) (string error)
}

type Manager struct{
	secretKey string
}

func NewManager(secretKey string)(*Manager, error){
	if secretKey == ""{
		return nil, ErrSecretKeyIsEmpty
	}
	return &Manager{secretKey}, nil
}

func (m *Manager) NewJWT(userID int, role string, ttl time.Duration) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": userID,
		"role": role,
		"exp": time.Now().Add(ttl).Unix(),
	})
	tokenString, err:= token.SignedString([]byte(m.secretKey))
	if err != nil{
		return "", err
	}
	return tokenString, nil
}

func (m *Manager) RefreshToken(oldToken string, refreshTokenTTL time.Duration) (string, error){
	token, err:= jwt.Parse(oldToken, func(token *jwt.Token)(interface{}, error){
		if _, ok:= token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(m.secretKey), nil
	})
	if err != nil{
		return "", err
	}

	if claims, ok:= token.Claims.(jwt.MapClaims); ok && token.Valid{
		userID := int(claims["id"].(float64))
		role := claims["role"].(string)
		return m.NewJWT(userID, role, refreshTokenTTL)
	}
	return "", ErrUnauthorized
}