package models

type Question struct {
	Id      uint   `json:"id"`
	Test_id uint   `json:"test_id"`
	Text    string `json:"text"`
}
