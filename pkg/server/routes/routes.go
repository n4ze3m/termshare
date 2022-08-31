package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/n4ze3m/termshare/pkg/server/handlers"
	"github.com/n4ze3m/termshare/pkg/server/middlewares"
)

var handler *handlers.Handler
var middleware *middlewares.Middleware

func init() {
	handler = handlers.NewHandler()
	middleware = middlewares.NewMiddleware()
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Middleware  []func(http.Handler) http.Handler
}

type Routes []Route

func NewRouter() *mux.Router {

	routes := Routes{
		Route{
			"Health",
			"GET",
			"/health",
			handler.HealthHandler,
			[]func(http.Handler) http.Handler{},
		},
		Route{
			"Websocket",
			"GET",
			"/ws",
			handler.WebsocketHandler,
			[]func(http.Handler) http.Handler{},
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Use(middleware.CORS)
		if len(route.Middleware) > 0 {
			for _, mw := range route.Middleware {
				router.
					Methods(route.Method).
					Path(route.Pattern).
					Name(route.Name).
					Handler(mw(route.HandlerFunc))
			}
		} else {
			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(route.HandlerFunc)
		}
	}
	return router
}
