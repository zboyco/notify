package middleware

import "net/http"

type AuthMiddleware struct {
	token string
}

func NewAuthMiddleware(token string) *AuthMiddleware {
	return &AuthMiddleware{
		token: token,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		// 判断请求头中是否包含token
		if r.Header.Get("token") != m.token {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// Passthrough to next handler if need
		next(w, r)
	}
}
