package models

import (
	"encoding/json"
)

type PortofolioList struct {
	ID          int    `json:"id" form:"id" xml:"id" example:"1"`
	Title       string `json:"title" form:"title" xml:"title" example:"IPKD"`
	Description string `json:"description" form:"description" xml:"description" example:"Indeks Pengelolaan Keuangan Daerah"`
	Images      string `json:"images" form:"images" xml:"images" example:"/assets/*png"`
}

func (p *PortofolioList) FromJSON(msg []byte) error {
	return json.Unmarshal(msg, p)
}

func (p *PortofolioList) ToJSON() []byte {
	str, _ := json.Marshal(p)
	return str
}
