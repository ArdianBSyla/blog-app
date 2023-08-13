package middleware

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
)

func JSONHeader() func(next http.Handler) http.Handler {
	return middleware.SetHeader("Content-Type", "application/json")
}
