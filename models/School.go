package models

import "time"

type School struct {
	Id                      uint      `json:"id"`
	Name                    string    `json:"name"`
	Location                string    `json:"location"`
	License_expiration_date time.Time `json:"license_expiration_date"`
}
