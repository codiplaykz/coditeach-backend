package models

type Subject struct {
	Id          uint   `json:"id"`
	Teacher_id  uint   `json:"teacher_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
