package models

import (
	"encoding/json"
)

type Environment struct {
	BASE_SERVER_NAME         string
	BASE_SERVER_URL          string
	BASE_SERVER_READ_TIMEOUT int
	BASE_DB_SERVER_URL       string
}

func (p *Environment) FromJSON(msg []byte) error {
	return json.Unmarshal(msg, p)
}

func (p *Environment) ToJSON() []byte {
	str, _ := json.Marshal(p)
	return str
}
