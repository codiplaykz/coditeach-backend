package models

import "time"

type CurriculumLesson struct {
	Id          uint      `json:"id"`
	Block_id    uint      `json:"block_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Duration    int       `json:"duration"`
	Content     string    `json:"content"`
	Created_at  time.Time `json:"created_at"`
}
