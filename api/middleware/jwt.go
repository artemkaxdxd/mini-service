package middleware

import (
	"context"
	"net/http"
	"strings"

	utils "github.com/artemkaxdxd/mini-service"
)

func ValidateToken(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("invalid token"))
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]

		id, err := utils.ValidateJWT(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		ctx := context.WithValue(r.Context(), "userId", id)

		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}
