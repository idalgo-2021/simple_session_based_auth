package controllers

import (
	"encoding/json"
	"net/http"
	"simple_session_based_auth/models"
	"simple_session_based_auth/repo"
)

func LogIn(w http.ResponseWriter, r *http.Request) {

	var creds models.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, exists := repo.GetUser(creds.Username)
	if !exists || user.Password != creds.Password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	session := repo.CreateSession(user.ID)
	setSessionCookie(w, session)
}

func setSessionCookie(w http.ResponseWriter, session models.Session) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    session.SessionID,
		Expires:  session.Expiration,
		HttpOnly: true,
	})
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	sessionCookie, err := r.Cookie("session_id")
	if err != nil || sessionCookie == nil {
		http.Error(w, "No session found", http.StatusUnauthorized)
		return
	}

	session, exists := repo.GetSession(sessionCookie.Value)

	if !exists {
		http.Error(w, "Unauthorized: Session not found", http.StatusUnauthorized)
		return
	}

	repo.DeleteSession(session.SessionID)

	newSession := repo.CreateSession(session.UserID)
	setSessionCookie(w, newSession)

}

func LogOut(w http.ResponseWriter, r *http.Request) {

	sessionCookie, err := r.Cookie("session_id")
	if err != nil || sessionCookie == nil {
		http.Error(w, "No session found", http.StatusUnauthorized)
		return
	}

	repo.DeleteSession(sessionCookie.Value)

	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		MaxAge: -1,
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logged out successfully"))
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	sessionCookie, err := r.Cookie("session_id")
	if err != nil || sessionCookie == nil {
		http.Error(w, "No session found", http.StatusUnauthorized)
		return
	}

	session, exists := repo.GetSession(sessionCookie.Value)

	if !exists {
		http.Error(w, "Unauthorized: Session not found", http.StatusUnauthorized)
		return
	}

	userID := session.UserID
	sessions := repo.GetSessionsForUser(userID)

	response, err := json.Marshal(sessions)
	if err != nil {
		http.Error(w, "Failed to marshal sessions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}
