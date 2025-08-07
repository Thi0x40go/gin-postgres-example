package models

type Widget struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"size:100;not null"`
	Weight int64
}
