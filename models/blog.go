package models

import "time"

type Blog struct {
	ID               uint
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Title            string
	ShortDescription string
	Description      string
	Published        bool
}
