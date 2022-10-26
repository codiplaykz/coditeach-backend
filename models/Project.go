package models

type Project struct {
	Id              uint   `json:"id"`
	Name            string `json:"name"`
	Type            string `json:"type"`
	Level           string `json:"level"`
	Description     string `json:"description"`
	Tech_components string `json:"tech_components"`
	Creator_Id      uint   `json:"creator_id"`
	Source_Code     string `json:"source_code"`
	Block_code      string `json:"block_code"`
	Cover_img_url   string `json:"cover_img_url"`
	Scheme_img_url  string `json:"scheme_img_url"`
}
