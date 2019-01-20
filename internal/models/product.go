package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	ProductName        string
	ProductPrice       float64
	ProductDescription string `gorm:"type:text;"`
	ProductImage       string
}