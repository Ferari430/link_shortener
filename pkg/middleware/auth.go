package middleware

import (
	"context"
	"log"
	"my_project/configs"
	"my_project/pkg/jwt"
	"net/http"
	"strings"
)

type key string

const (
	ContextEmailKey key = "ContextEmailKey"
)

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		token := parts[1]
		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		if !isValid {
			http.Error(w, "Invalid JWT token", 402)
			return
		}

		ctx := context.WithValue(r.Context(), ContextEmailKey, data.Email)

		req := r.WithContext(ctx)

		log.Println("AUTHORIZATION SUCCSESS By IsAuthed ")

		log.Println(isValid)
		log.Println(data)

		log.Println("Received AUTH token:", token)

		next.ServeHTTP(w, req)
	})
}
