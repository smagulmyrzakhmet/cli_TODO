package models

import "time"

type Task struct {
	Id          uint
	Title       string
	Description string
	Status      Status
	CreatedAt   time.Time
}
