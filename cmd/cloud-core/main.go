package main

import (
	"net/http"

	"github.com/RiddlerXenon/cloud-core/internal/middleware"
	"github.com/RiddlerXenon/cloud-core/internal/routes"
	"go.uber.org/zap"
)

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
}

func main() {
	mux := http.NewServeMux()

	routes := routes.RegisterRoutes(mux)

	handlerWithMiddleware := middleware.CORS(mux)

	zap.S().Info("Starting server on :8080")
	if err := http.ListenAndServe(":8080", handlerWithMiddleware); err != nil {
		zap.S().Fatal("Failed to start server: ", err)
	}
}
