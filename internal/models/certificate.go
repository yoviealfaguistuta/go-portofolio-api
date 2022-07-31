package models

import (
	"encoding/json"
)

type Certificate struct {
	ID          int    `json:"id" form:"id" xml:"id" example:"1"`
	Images      string `json:"images" form:"images" xml:"images" example:"assets/images/certificate/*.png"`
	Title       string `json:"title" form:"title" xml:"title" example:"Android Developer Roadmap"`
	Publish     string `json:"publish" form:"publish" xml:"publish" example:"Agustus 2021"`
	Credentials string `json:"credentials" form:"credentials" xml:"credentials" example:"QLZ9144K7P5D"`
	Urls        string `json:"urls" form:"urls" xml:"urls" example:"https://www.dicoding.com/certificates/*"`
}

func (p *Certificate) FromJSON(msg []byte) error {
	return json.Unmarshal(msg, p)
}

func (p *Certificate) ToJSON() []byte {
	str, _ := json.Marshal(p)
	return str
}
