package middlewares

import (
	"log"
	"net/http"
)

// This will recover application is panic happens in the application
func RcoverMiddlerware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Recovered Please check the error:-", err)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
