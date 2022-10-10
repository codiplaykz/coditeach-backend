package models

import (
	"time"
)

type Block struct {
	Id         uint      `json:"id"`
	Module_id  uint      `json:"module_id"`
	Title      string    `json:"title"`
	Created_at time.Time `json:"created_at"`
}
