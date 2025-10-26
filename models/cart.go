package models

type Cart struct {
	Id         uint `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	TotalPrice uint `gorm:"not null" json:"total_price"`

	UserId uint  `gorm:"not null; uniqueIndex" json:"user_id"`
	User   *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`

	Products []Product `gorm:"many2many:cart_products" json:"products"`
}
