package models

type Product struct {
	Id          uint    `gorm:"primary_key;auto_increment" json:"id"`
	Title       string  `gorm:"not null" json:"title"`
	Description string  `gorm:"not null" json:"description"`
	Price       float32 `gorm:"not null" json:"price"`
	CreatedBy   uint    `gorm:"not null;" json:"user_id"`
}
