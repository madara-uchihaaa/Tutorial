package middleware

import (
	"backend/utils"
	"net/http"
)

func VerifyToken(originalHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if !utils.ValidateToken(token) {
			panic("Invalid Token")
		}
		originalHandler.ServeHTTP(w, r)
	})
}
