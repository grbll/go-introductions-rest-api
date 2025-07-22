package models

import "time"

type UserSession struct {
	UserID       int
	UserMail     string
	SessionStart *time.Time
	SessionEnd   *time.Time
}
