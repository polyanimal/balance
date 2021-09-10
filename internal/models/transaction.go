package models

import "time"

type Transaction struct {
	Id         int    `json:"id"`
	Comment    string    `json:"comment"`
	IdFrom     string    `json:"id_from"`
	IdTo       string    `json:"id_to"`
	DateCreate time.Time `json:"creation_date"`
	Funds      int       `json:"funds"`
}
