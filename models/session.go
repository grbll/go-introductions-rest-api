package models

import "time"

type ActiveSession struct {
	SessionId    int
	UserId       int
	SessionStart *time.Time
}

type ClosedSession struct {
	ActiveSession
	SessionEnd *time.Time
}
