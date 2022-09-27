package models

type Class struct {
	Id         uint   `json:"id"`
	Teacher_id uint   `json:"teacher_id"`
	School_id  uint   `json:"school_id"`
	Name       string `json:"name"`
	Code       string `json:"code"`
}
