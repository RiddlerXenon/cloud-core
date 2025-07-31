package main

import (
	"net/http"

	"github.com/RiddlerXenon/cloud-core/internal/config"
	"github.com/RiddlerXenon/cloud-core/internal/handlers"
	"github.com/RiddlerXenon/cloud-core/internal/middleware"
	"github.com/RiddlerXenon/cloud-core/internal/repository"
	"github.com/RiddlerXenon/cloud-core/internal/routes"
	"go.uber.org/zap"
)

func main() {
	// Initialize the logger
	logger := zap.Must(zap.NewProduction())
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	// Initialize the configuration
	cfg, err := config.InitConfig()
	if err != nil {
		zap.S().Fatal("Failed to initialize configuration: ", err)
	}

	d, err := repository.InitDB(cfg)
	if err != nil {
		zap.S().Fatal("Failed to initialize DB: ", err)
	}
	defer d.DB.Close()

	h := handlers.New(d, cfg)

	mux := http.NewServeMux()
	routes.RegisterRoutes(mux, h)
	handlerWithMiddleware := middleware.CORS(mux)

	zap.S().Info("Starting server on :8080")
	if err := http.ListenAndServe(":8080", handlerWithMiddleware); err != nil {
		zap.S().Fatal("Failed to start server: ", err)
	}
}
