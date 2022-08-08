package models

import (
	"encoding/json"
)

type PortofolioList struct {
	ID           int         `json:"id" form:"id" xml:"id" example:"1"`
	Title        string      `json:"title" form:"title" xml:"title" example:"IPKD"`
	Descriptions string      `json:"descriptions" form:"descriptions" xml:"descriptions" example:"Indeks Pengelolaan Keuangan Daerah"`
	CreatedAt    interface{} `json:"updated_at" form:"updated_at" xml:"created_at"`
	Images       string      `json:"images" form:"images" xml:"images" example:"/assets/*png"`
}

func (p *PortofolioList) FromJSON(msg []byte) error {
	return json.Unmarshal(msg, p)
}

func (p *PortofolioList) ToJSON() []byte {
	str, _ := json.Marshal(p)
	return str
}
