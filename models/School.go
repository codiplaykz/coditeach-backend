package models

type School struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Code     string `json:"code"`
}
