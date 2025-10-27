package models

type User struct {
	Id       uint   `gorm:"primary_key;auto_increment" json:"id"`
	Username string `gorm:"unique; not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Role     string `gorm:"not null;default:'customer'" json:"role"`
}
