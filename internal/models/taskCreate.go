package models

import "time"

type TaskCreate struct {
	Title       string
	Description string
	CreateAt    time.Time
}
