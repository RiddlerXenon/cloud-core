package auth

import (
	"fmt"
	"time"

	"github.com/RiddlerXenon/cloud-core/internal/config"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

func GenerateJWT(cfg *config.Config, userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * time.Duration(cfg.JWTExpire)).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if cfg.JWTSecret == "" {
		return "", fmt.Errorf("JWT secret is not set")
	}

	return token.SignedString([]byte(cfg.JWTSecret))
}

func ValidateJWT(tokenString string) (map[string]interface{}, error) {
	cfg := config.GetConfig()

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		if cfg.JWTSecret == "" {
			return nil, fmt.Errorf("JWT secret is not set")
		}

		return []byte(cfg.JWTSecret), nil
	})
	if err != nil {
		zap.S().Errorw("Error parsing token", "error", err, "token", tokenString)
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		zap.S().Errorw("Invalid token", "token", tokenString)
		return nil, fmt.Errorf("Invalid token")
	}
}
