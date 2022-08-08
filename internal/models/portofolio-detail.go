package models

import (
	"encoding/json"
)

type PortofolioDetail struct {
	ID           int                     `json:"id" form:"id" xml:"id" example:"1"`
	Title        string                  `json:"title" form:"title" xml:"title" example:"IPKD"`
	Descriptions string                  `json:"descriptions" form:"descriptions" xml:"descriptions" example:"Indeks Pengelolaan Keuangan Daerah"`
	ProjectInfo  string                  `json:"project_info" form:"project_info" xml:"project_info" example:"Sistem ini dibangun pada tahun 2021"`
	Languages    string                  `json:"languages" form:"languages" xml:"languages" example:"Javascript"`
	Tech         string                  `json:"tech" form:"tech" xml:"tech" example:"React JS"`
	JobRole      string                  `json:"job_role" form:"job_role" xml:"job_role" example:"Server Side"`
	Databases    string                  `json:"databases" form:"databases" xml:"databases" example:"PostgreSql"`
	Dates        string                  `json:"dates" form:"dates" xml:"dates" example:"1 Januari 2022"`
	Platform     string                  `json:"platform" form:"platform" xml:"platform" example:"Website"`
	Urls         string                  `json:"urls" form:"urls" xml:"urls" example:"https://ipkd-bpp.kemendagri.go.id"`
	SourceCode   string                  `json:"source_code" form:"source_code" xml:"source_code" example:"https://github.com/*project"`
	CreatedAt    interface{}             `json:"updated_at" form:"updated_at" xml:"updated_at"`
	UpdatedAt    interface{}             `json:"created_at" form:"created_at" xml:"created_at"`
	Images       []PortofolioImagesModel `json:"images" form:"images" xml:"images"`
}

type PortofolioImagesModel struct {
	Orders int64  `json:"orders" form:"orders" xml:"orders" example:"1"`
	Images string `json:"images" form:"images" xml:"images" example:"/assets/png"`
}

func (p *PortofolioDetail) FromJSON(msg []byte) error {
	return json.Unmarshal(msg, p)
}

func (p *PortofolioDetail) ToJSON() []byte {
	str, _ := json.Marshal(p)
	return str
}
