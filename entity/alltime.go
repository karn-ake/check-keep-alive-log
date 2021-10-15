package entity

import "time"

type Alert struct {
	Client string
	Status bool
}

type AllTime struct {
	LogTime    time.Time
	SystemTime time.Time
	DiffTime   time.Duration
}
