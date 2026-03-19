package middleware

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/golang-jwt/jwt/v5"
)

const defaultJWTSecret = "SUPER_SECRET_JWT_KEY_2026"

func JWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = defaultJWTSecret
	}
	return []byte(secret)
}

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			http.Error(w, "token não fornecido", http.StatusUnauthorized)
			return
		}
		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			http.Error(w, "header Authorization inválido", http.StatusUnauthorized)
			return
		}
		_, err := jwt.Parse(parts[1], func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("algoritmo de assinatura inválido")
			}
			return JWTSecret(), nil
		})
		if err != nil {
			http.Error(w, "token inválido", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func RequestID(next http.Handler) http.Handler { return middleware.RequestID(next) }
func RealIP(next http.Handler) http.Handler    { return middleware.RealIP(next) }
func Logger(next http.Handler) http.Handler    { return middleware.Logger(next) }
func Recoverer(next http.Handler) http.Handler { return middleware.Recoverer(next) }
