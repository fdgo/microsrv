package model

type Model struct {
	ID           int `gorm:"primary_key" json:"id"`
	Createded_At int `json:"created_at"`
	Updated_At   int `json:"updated_at"`
	IsDelete     int `json:"is_delete"`
}
