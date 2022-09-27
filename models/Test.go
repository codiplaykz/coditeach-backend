package models

import "time"

type Test struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Duration    int       `json:"duration"`
	Created_at  time.Time `json:"created_at"`
	Teacher_id  uint      `json:"teacher_id"`
}
