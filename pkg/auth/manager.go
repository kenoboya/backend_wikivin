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

	ErrInvalidUserID = errors.New("invalid user ID in token")
	ErrInvalidRole = errors.New("invalid role in token")
)

type TokenManager interface {
	NewJWT(userID int, role string, ttl time.Duration) (string, error)
	RefreshToken(token string, refreshTokenTTL time.Duration) (string, error)
	ParseToken(token string) (*Claims, error)
}

type Manager struct{
	secretKey string
}

type Claims struct {
	UserID int    `json:"id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
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

func (m *Manager) RefreshToken(oldToken string, refreshTokenTTL time.Duration) (string, error) {
	token, err := jwt.Parse(oldToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(m.secretKey), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["id"].(float64) 
		if !ok {
			return "", ErrInvalidUserID
		}
		role, ok := claims["role"].(string)
		if !ok {
			return "", ErrInvalidRole
		}

		newToken, err := m.NewJWT(int(userID), role, refreshTokenTTL)
		if err != nil {
			return "", err
		}

		return newToken, nil
	}
	return "", ErrUnauthorized
}

func (m *Manager) ParseToken(tokenString string) (*Claims, error){
	token, err:= jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token)(interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(m.secretKey), nil
	})
	if err != nil{
		return nil, err
	}
	if claims, ok:= token.Claims.(*Claims); ok && token.Valid{
		return claims, nil
	}
	return nil, ErrUnauthorized
}
