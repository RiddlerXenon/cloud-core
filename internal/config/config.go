package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Config struct {
	JWTSecret string
	JWTExpire int
	Username  string
	Password  string
	DBHost    string
	DBPort    int
	DBUser    string
	DBPass    string
	DBName    string
}

func InitConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		zap.S().Warn("Ошибка загрузки .env файла, используем переменные окружения")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		initErr := fmt.Errorf("JWT_SECRET не установлен")
		return nil, initErr
	}

	jwtExpireStr := os.Getenv("JWT_EXPIRE")
	if jwtExpireStr == "" {
		initErr := fmt.Errorf("JWT_EXPIRE не установлен")
		return nil, initErr
	}

	jwtExpire, err := strconv.Atoi(jwtExpireStr)
	if err != nil {
		initErr := fmt.Errorf("JWT_EXPIRE должен быть числом: %w", err)
		return nil, initErr
	}

	filePath := os.Getenv("CONFIG_FILE")
	if filePath == "" {
		initErr := fmt.Errorf("CONFIG_FILE не установлен")
		return nil, initErr
	}

	username, password, err := readCredentialsFromJSON(filePath)
	if err != nil {
		zap.S().Error("Ошибка чтения учетных данных из JSON файла", zap.Error(err))
		initErr := fmt.Errorf("ошибка чтения учетных данных из JSON файла: %w", err)
		return nil, initErr
	}

	dbHost := os.Getenv("DBHOST")
	if jwtExpireStr == "" {
		initErr := fmt.Errorf("DBHOST не установлен")
		return nil, initErr
	}

	dbPortStr := os.Getenv("DBPORT")
	if jwtExpireStr == "" {
		initErr := fmt.Errorf("DBPORT не установлен")
		return nil, initErr
	}

	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		initErr := fmt.Errorf("DBPORT должен быть числом: %w", err)
		return nil, initErr
	}

	dbUser := os.Getenv("DBUSER")
	if jwtExpireStr == "" {
		initErr := fmt.Errorf("DBUSER не установлен")
		return nil, initErr
	}

	dbPass := os.Getenv("DBPASS")
	if jwtExpireStr == "" {
		initErr := fmt.Errorf("DBPASS не установлен")
		return nil, initErr
	}

	dbName := os.Getenv("DBNAME")
	if jwtExpireStr == "" {
		initErr := fmt.Errorf("DBNAME не установлен")
		return nil, initErr
	}

	cfg = &Config{
		JWTSecret: jwtSecret,
		JWTExpire: jwtExpire,
		Username:  username,
		Password:  password,
		DBHost:    dbHost,
		DBPort:    dbPort,
		DBUser:    dbUser,
		DBPass:    dbPass,
		DBName:    dbName,
	}

	if cfg.JWTSecret == "" || cfg.JWTExpire <= 0 || cfg.Username == "" || cfg.Password == "" {
		initErr := fmt.Errorf("конфигурация содержит некорректные данные")
		return nil, initErr
	}

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
