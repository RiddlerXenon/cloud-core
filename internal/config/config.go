package config

import (
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

		cfg = &Config{
			JWTSecret: jwtSecret,
			JWTExpire: jwtExpire,
		}
	})

	return cfg, initErr
}

func GetConfig() *Config {
	if cfg == nil {
		zap.S().Error("Конфигурация не инициализирована.")
	}
	return cfg
}
