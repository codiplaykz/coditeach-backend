package models

type Option struct {
	Id          uint   `json:"id"`
	Question_id uint   `json:"question_id"`
	Text        string `json:"option_text"`
	Is_correct  bool   `json:"is_correct"`
}
