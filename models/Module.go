package models

import "time"

type Module struct {
	Id            uint      `json:"id"`
	Curriculum_id uint      `json:"curriculum_id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Created_at    time.Time `json:"created_at"`
}
