package models

type Posts struct {
	Id          int    `gorm:"type:int;primary_key"`
	Title       string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:varchar(255)"`
}
