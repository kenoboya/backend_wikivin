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

	AccessToken = "access"
	RefreshToken = "refresh"
)

type TokenManager interface {
	NewJWT(userID int, role string, ttl time.Duration, tokenType string) (string, error)
	RefreshToken(token string, refreshTokenTTL time.Duration) (string, error)
	ParseToken(token string, tokenType string) (*Claims, error)
	VerifyToken(token string, tokenType string) error
}

type Manager struct{
	secretAccessKey string
	secretRefreshKey string
}

type Claims struct {
	UserID int    `json:"id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func NewManager(secretAccessKey string, secretRefreshKey string)(*Manager, error){
	if secretAccessKey == "" && secretRefreshKey == ""{
		return nil, ErrSecretKeyIsEmpty
	}
	return &Manager{
		secretAccessKey, 
		secretRefreshKey,
		}, nil
}

func (m *Manager) NewJWT(userID int, role string, ttl time.Duration, tokenType string) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": userID,
		"role": role,
		"exp": time.Now().Add(ttl).Unix(),
	})
	var secretKey []byte
	switch tokenType {
	case AccessToken:
		secretKey = []byte(m.secretAccessKey)
	case RefreshToken:
		secretKey = []byte(m.secretAccessKey + m.secretRefreshKey)
	default:
		return "", fmt.Errorf("unknown token type: %s", tokenType)
	}
	tokenString, err:= token.SignedString(secretKey)
	if err != nil{
		return "", err
	}
	return tokenString, nil
}

func (m *Manager) VerifyToken(token string, tokenType string) error {
	var secretKey []byte

	switch tokenType {
	case AccessToken:
		secretKey = []byte(m.secretAccessKey)
	case RefreshToken:
		secretKey = []byte(m.secretAccessKey + m.secretRefreshKey)
	default:
		return fmt.Errorf("unknown token type: %s", tokenType)
	}

	parsedToken, err := jwt.Parse(token, func(parsedToken *jwt.Token) (interface{}, error) {
		if _, ok := parsedToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", parsedToken.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		if _, ok := claims["id"].(float64); !ok {
			return ErrInvalidUserID
		}
		if _, ok := claims["role"].(string); !ok {
			return ErrInvalidRole
		}
		return nil
	}
	return ErrUnauthorized
}

func (m *Manager) RefreshToken(oldToken string, refreshTokenTTL time.Duration) (string, error) {
	token, err := jwt.Parse(oldToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(m.secretAccessKey + m.secretRefreshKey), nil
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

		newToken, err := m.NewJWT(int(userID), role, refreshTokenTTL, RefreshToken)
		if err != nil {
			return "", err
		}

		return newToken, nil
	}
	return "", ErrUnauthorized
}

func (m *Manager) ParseToken(tokenString string, tokenType string) (*Claims, error){
	token, err:= jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token)(interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		var secretKey []byte
		switch tokenType {
		case AccessToken:
			secretKey = []byte(m.secretAccessKey)
		case RefreshToken:
			secretKey = []byte(m.secretAccessKey + m.secretRefreshKey)
		default:
			return nil, fmt.Errorf("unknown token type: %s", tokenType)
		}
		return []byte(secretKey), nil
	})
	if err != nil{
		return nil, err
	}
	if claims, ok:= token.Claims.(*Claims); ok && token.Valid{
		return claims, nil
	}
	return nil, ErrUnauthorized
}
