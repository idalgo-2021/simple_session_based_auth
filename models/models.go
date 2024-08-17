package models

import "time"

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID       int
	Username string
	Password string
}

type Session struct {
	SessionID  string
	UserID     int
	Expiration time.Time
}
