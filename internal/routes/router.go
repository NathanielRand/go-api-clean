package routes

import (
	"github.com/NathanielRand/go-api-clean/internal/handlers"
	"github.com/NathanielRand/go-api-clean/internal/middleware"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func SetupRouter() *mux.Router {
	// Initialize a new router from the Gorilla Mux library
	router := mux.NewRouter()

	// Create a new middleware chain using the Alice library
	chain := alice.New()

	// Middleware chain for authentication, rate limiting, caching, and quotas
	// chain = chain.Append(middleware.SecurityMiddleware)
	// chain = chain.Append(middleware.AuthenticationMiddleware)
	// chain = chain.Append(func(next http.Handler) http.Handler {
	// 	return middleware.AuthorizationMiddleware(next, "admin")
	// })
	// chain = chain.Append(middleware.RateLimitingMiddleware)
	// chain = chain.Append(middleware.QuotaMiddleware)
	// chain = chain.Append(middleware.CachingMiddleware)
	chain = chain.Append(middleware.LoggingMiddleware)

	// Health endpoints
	router.Handle("/api/v1/health", chain.ThenFunc(handlers.HealthHandler)).Methods("GET")

	// Return the router
	return router
}
