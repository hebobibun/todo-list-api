package activity

import "time"

type Core struct {
	ID        uint
	Title     string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
