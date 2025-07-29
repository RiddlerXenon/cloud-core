package auth

import (
	"fmt"

	"github.com/RiddlerXenon/cloud-core/internal/config"
	"go.uber.org/zap"
)

func Login(u, p string, cfg *config.Config) (string, error) {
	username := cfg.Username
	password := cfg.Password

	if u != username || p != password {
		zap.S().Errorw("Неверные учетные данные", "username", u, "password", p)
		return "", fmt.Errorf("неверные учетные данные")
	}

	token, err := GenerateJWT(cfg, username)
	if err != nil {
		zap.S().Errorw("Ошибка генерации JWT", "error", err)
		return "", err
	}

	return token, nil
}
