package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Config struct {
	JWTSecret string
	JWTExpire int
	Username  string
	Password  string
}

var (
	cfg  *Config
	once sync.Once
)

func InitConfig() (*Config, error) {
	var initErr error

	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			zap.S().Warn("Ошибка загрузки .env файла, используем переменные окружения")
		}

		jwtSecret := os.Getenv("JWT_SECRET")
		if jwtSecret == "" {
			zap.S().Error("JWT_SECRET не установлен в переменных окружения")
			initErr = fmt.Errorf("JWT_SECRET не установлен")
			return
		}

		jwtExpireStr := os.Getenv("JWT_EXPIRE")
		if jwtExpireStr == "" {
			zap.S().Error("JWT_EXPIRE не установлен в переменных окружения")
			initErr = fmt.Errorf("JWT_EXPIRE не установлен")
			return
		}

		jwtExpire, err := strconv.Atoi(jwtExpireStr)
		if err != nil {
			zap.S().Error("JWT_EXPIRE должен быть числом", zap.Error(err))
			initErr = fmt.Errorf("JWT_EXPIRE должен быть числом: %w", err)
			return
		}

		filePath := os.Getenv("CONFIG_FILE")
		if filePath == "" {
			zap.S().Error("CONFIG_FILE не установлен в переменных окружения")
			initErr = fmt.Errorf("CONFIG_FILE не установлен")
			return
		}

		username, password, err := readCredentialsFromJSON(filePath)
		if err != nil {
			zap.S().Error("Ошибка чтения учетных данных из JSON файла", zap.Error(err))
			initErr = fmt.Errorf("ошибка чтения учетных данных из JSON файла: %w", err)
			return
		}

		cfg = &Config{
			JWTSecret: jwtSecret,
			JWTExpire: jwtExpire,
			Username:  username,
			Password:  password,
		}

		if cfg.JWTSecret == "" || cfg.JWTExpire <= 0 || cfg.Username == "" || cfg.Password == "" {
			zap.S().Error("Конфигурация содержит некорректные данные")
			initErr = fmt.Errorf("конфигурация содержит некорректные данные")
			return
		}
	})

	return cfg, initErr
}

func readCredentialsFromJSON(filePath string) (string, string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		zap.S().Error("Ошибка открытия файла", zap.String("filePath", filePath), zap.Error(err))
		return "", "", fmt.Errorf("ошибка открытия файла %s: %w", filePath, err)
	}

	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err = json.Unmarshal(data, &credentials)
	if err != nil {
		zap.S().Error("Ошибка декодирования JSON", zap.String("filePath", filePath), zap.Error(err))
		return "", "", fmt.Errorf("ошибка декодирования JSON: %w", err)
	}

	if credentials.Username == "" || credentials.Password == "" {
		zap.S().Error("Пустые учетные данные в JSON", zap.String("filePath", filePath))
		return "", "", fmt.Errorf("пустые учетные данные в JSON")
	}
	return credentials.Username, credentials.Password, nil
}
