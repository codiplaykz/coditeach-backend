package models

import "time"

type Curriculum struct {
	Id          uint      `json:"id"`
	User_id     uint      `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Created_at  time.Time `json:"created_at"`
}
