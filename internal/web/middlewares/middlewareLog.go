package middlewares

import (
	"log"
	"net/http"
	"time"
)

func MiddlewareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Início %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Término %s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}
