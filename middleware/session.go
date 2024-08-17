package middleware

import (
	"net/http"
	"simple_session_based_auth/models"
	"simple_session_based_auth/repo"
	"time"
)

func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/login" || r.URL.Path == "/logout" {
			next.ServeHTTP(w, r)
			return
		}

		sessionCookie, err := r.Cookie("session_id")
		if err != nil || sessionCookie == nil {
			http.Error(w, "Unauthorized: No session found", http.StatusUnauthorized)
			return
		}

		session, exists := repo.GetSession(sessionCookie.Value)
		if !exists || IsSessionExpired(session) {
			http.Error(w, "Unauthorized: Session expired or not found", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func IsSessionExpired(session models.Session) bool {
	return time.Now().After(session.Expiration)
}
