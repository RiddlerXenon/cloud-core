package auth

import (
	"fmt"

	"github.com/RiddlerXenon/internal/config"
	"go.uber.org/zap"
)

func Login(u, p string) (string, error) {
	username := config.GetConfig().Username
	password := config.GetConfig().Password

	if u != username || p != password {
		zap.S().Errorw("Неверные учетные данные", "username", u, "password", p)
		return "", fmt.Errorf("неверные учетные данные")
	}

	token, err := GenerateJWT(username)
	if err != nil {
		zap.S().Errorw("Ошибка генерации JWT", "error", err)
		return "", err
	}

	return token, nil
}
