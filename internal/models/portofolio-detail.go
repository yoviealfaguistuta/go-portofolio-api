package models

import (
	"encoding/json"
)

type PortofolioDetail struct {
	ID          int    `json:"id" form:"id" xml:"id" example:"1"`
	Title       string `json:"title" form:"title" xml:"title" example:"IPKD"`
	Description string `json:"description" form:"description" xml:"description" example:"description" example:"Indeks Pengelolaan Keuangan Daerah"`
	Images      string `json:"images" form:"images" xml:"images" example:"/assets/*png"`
	Database    string `json:"database" form:"database" xml:"database" example:"PostgreSql"`
	Dates       string `json:"dates" form:"dates" xml:"dates" example:"1 Januari 2022"`
	Languages   string `json:"languages" form:"languages" xml:"languages" example:"Javascript"`
	Platform    string `json:"platform" form:"platform" xml:"platform" example:"Website"`
	ProjectInfo string `json:"project_info" form:"project_info" xml:"project_info" example:"Sistem ini dibangun pada tahun 2021"`
	SourceCode  string `json:"source_code" form:"source_code" xml:"source_code" example:"https://github.com/*project"`
	Url         string `json:"url" form:"url" xml:"url" example:"https://ipkd-bpp.kemendagri.go.id"`
	CreatedAt   string `json:"updated_at" form:"updated_at" xml:"updated_at" example:"2021-08-03 18:00:19"`
	UpdatedAt   string `json:"created_at" form:"created_at" xml:"created_at" example:"2021-08-03 18:00:19"`
}

func (p *PortofolioDetail) FromJSON(msg []byte) error {
	return json.Unmarshal(msg, p)
}

func (p *PortofolioDetail) ToJSON() []byte {
	str, _ := json.Marshal(p)
	return str
}
