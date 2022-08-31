package routes

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/n4ze3m/termshare/pkg/server/handlers"
	"github.com/n4ze3m/termshare/pkg/server/middlewares"
)

var handler *handlers.Handler
var middleware *middlewares.Middleware
var ctx context.Context
var rdb *redis.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ctx = context.Background()
	option, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatal(err)
	}
	rdb = redis.NewClient(option)

	handler = handlers.NewHandler(&ctx, rdb)
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
