package models

import (
	"e-commerce/config"
	"log"
)

func Migrate() {
	if err := config.DB.AutoMigrate(&User{}, &Product{}, &Cart{}); err != nil {
		log.Fatal(err)
	}
}
