package models

type Schedule struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	Subject_id uint   `json:"subject_id"`
	Class_id   uint   `json:"class_id"`
}
