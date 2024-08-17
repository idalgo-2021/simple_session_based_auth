package repo

import (
	"simple_session_based_auth/models"
	"time"

	"github.com/google/uuid"
)

var Users = map[string]models.User{
	"user1": {ID: 1, Username: "user1", Password: "pass1"},
	"user2": {ID: 2, Username: "user2", Password: "pass2"},
}

var Sessions = map[string]models.Session{}

func GetUser(username string) (models.User, bool) {
	user, exists := Users[username]
	return user, exists
}

func generateSessionID() string {
	return uuid.NewString()
}

func CreateSession(userID int) models.Session {
	sessionID := generateSessionID()
	expiration := time.Now().Add(3 * time.Minute)

	session := models.Session{
		SessionID:  sessionID,
		UserID:     userID,
		Expiration: expiration,
	}

	SaveSession(sessionID, session)
	return session
}

func SaveSession(sessionID string, session models.Session) {
	Sessions[sessionID] = session
}

func GetSession(sessionID string) (models.Session, bool) {
	session, exists := Sessions[sessionID]
	return session, exists
}

func DeleteSession(sessionID string) {
	delete(Sessions, sessionID)
}

func GetSessionsForUser(userID int) []models.Session {
	sessions := []models.Session{}
	for _, session := range Sessions {
		if session.UserID == userID {
			sessions = append(sessions, session)
		}
	}
	return sessions
}
