package models

import (
	"time"
)

type Statistics struct {
	Id          uint      `json:"id"`
	Information string    `json:"information"`
	Created_at  time.Time `json:"created_at"`
}
