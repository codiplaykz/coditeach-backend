package models

import (
	"time"
)

type Homework struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Subject_id  uint      `json:"subject_id"`
}
