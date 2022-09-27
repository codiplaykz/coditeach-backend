package models

import (
	"time"
)

type Report struct {
	Id            uint      `json:"id"`
	Information   string    `json:"information"`
	Download_link string    `json:"download_link"`
	Created_at    time.Time `json:"created_at"`
}
