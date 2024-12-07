package middleware

import (
	"devbook-api/pkg/auth"
	"devbook-api/pkg/response"
	"net/http"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidateToken(r); err != nil {
			response.Error(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	})
}
