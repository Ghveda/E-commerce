package models

type Product struct {
	Id          uint   `gorm:"primary_key;auto_increment" json:"id"`
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"not null" json:"description"`
	Price       uint   `gorm:"not null" json:"price"`
}
