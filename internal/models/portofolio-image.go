package models

import (
	"encoding/json"
)

type PortofolioImage struct {
	Images string `json:"images" form:"images" xml:"images" example:"/assets/*png"`
	Orders int    `json:"orders" form:"orders" xml:"orders" example:"1"`
}

func (p *PortofolioImage) FromJSON(msg []byte) error {
	return json.Unmarshal(msg, p)
}

func (p *PortofolioImage) ToJSON() []byte {
	str, _ := json.Marshal(p)
	return str
}
