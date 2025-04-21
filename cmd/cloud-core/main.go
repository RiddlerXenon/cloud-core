package main

import (
	"net/http"

	"github.com/RiddlerXenon/cloud-core/internal/config"
	"github.com/RiddlerXenon/cloud-core/internal/middleware"
	"github.com/RiddlerXenon/cloud-core/internal/routes"
	"go.uber.org/zap"
)

func main() {
	// Initialize the logger
	logger := zap.Must(zap.NewProduction())
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	// Initialize the configuration
	_, err := config.InitConfig()
	if err != nil {
		zap.S().Fatal("Failed to initialize configuration: ", err)
	}

	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)
	handlerWithMiddleware := middleware.CORS(mux)

	zap.S().Info("Starting server on :8080")
	if err := http.ListenAndServe(":8080", handlerWithMiddleware); err != nil {
		zap.S().Fatal("Failed to start server: ", err)
	}
}
