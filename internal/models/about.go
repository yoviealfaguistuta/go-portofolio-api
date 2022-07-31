package models

import (
	"encoding/json"
)

type About struct {
	Message string `json:"message,omitempty" form:"message" xml:"message" example:"I am .. years old"`
}

func (p *About) FromJSON(msg []byte) error {
	return json.Unmarshal(msg, p)
}

func (p *About) ToJSON() []byte {
	str, _ := json.Marshal(p)
	return str
}
