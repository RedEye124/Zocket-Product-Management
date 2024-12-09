package models

// Product represents the product model
type Product struct {
	ID                      uint   `gorm:"primary_key"`
	UserID                  uint   `gorm:"not null"`
	ProductName             string `gorm:"size:255;not null"`
	ProductDescription      string
	ProductImages           []string
	ProductPrice            float64 `gorm:"not null"`
	CompressedProductImages []string
}
