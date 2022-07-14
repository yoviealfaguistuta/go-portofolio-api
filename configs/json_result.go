package configs

type JSONResult struct {
	Code    int         `json:"code" example:"404"`
	Message string      `json:"message" example:"Not Found"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta"`
}

type JSONResultMeta struct {
	TotalCount  int `json:"total_count"`
	PageCount   int `json:"page_count"`
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
}

type JSONResultLogin struct {
	Token        string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzYwODU0MjksImlkIjoyLCJwaG9uZSI6Iis2MjgxMjM0NTYyIiwidXNlcm5hbWUiOi.dl_ojy9ojLnWqpW589YltLPV61TCsON-3yQ2"`
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzYxNjgyMjksImlk._aYI7pV2c9SU9VOp3RY_mxtFenYFQuKPJtVfk"`
}
