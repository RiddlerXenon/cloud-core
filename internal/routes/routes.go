package routes

import (
	"net/http"

	"github.com/RiddlerXenon/cloud-core/internal/handlers"
)

type Route struct {
	Path    string
	Handler http.HandlerFunc
}

func RegisterRoutes(mux *http.ServeMux) {
	routes := []Route{
		{Path: "/auth/login", Handler: handlers.LoginHandler},
		{Path: "/auth/logout", Handler: handlers.LogoutHandler},
	}

	for _, route := range routes {
		mux.HandleFunc(route.Path, route.Handler)
	}
}
