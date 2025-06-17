package models

import "time"

type UserSession struct {
	UserID       int
	UserName     string
	SessionStart *time.Time
	SessionEnd   *time.Time
}
