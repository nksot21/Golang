package models

import "time"

type Order struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	ProductRefer int     `json:"productid"`
	Product      Product `gorm:"foreignKey:ProductRefer"`
	UserRefer    int     `json:"userid"`
	User         User    `gorm:"foreignKey:UserRefer"`
}
