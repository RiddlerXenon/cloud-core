// package auth

// import (
// 	"fmt"

// 	"github.com/RiddlerXenon/cloud-core/internal/config"
// 	"go.uber.org/zap"
// )

// func Login(u, p string, ) (string, error) {
// 	if !auth.VerifyPassword(p, hash) {
// 		return nil, fmt.Error("Password not verified!")
// 	}

// 	token, err := auth.GenerateJWT(d.Cfg, u)
// 	if err != nil {
// 		return nil, fmt.Errorf("JWT generation error: %e", err)
// 	}

// 	return token, nil
// }

// func Login(u, p string, cfg *config.Config) (string, error) {
// 	username := cfg.Username
// 	password := cfg.Password

// 	if u != username || p != password {
// 		zap.S().Errorw("Неверные учетные данные", "username", u, "password", p)
// 		return "", fmt.Errorf("неверные учетные данные")
// 	}

// 	token, err := GenerateJWT(cfg, username)
// 	if err != nil {
// 		zap.S().Errorw("Ошибка генерации JWT", "error", err)
// 		return "", err
// 	}

// 	return token, nil
// }
