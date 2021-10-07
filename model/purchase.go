package model

import "time"

type Purchase struct {
	UID 			string				`gorm:"primaryKey"`
	Item 			string
	Product 		Product				`gorm:"foreignKey:Item"`
	Quantity		uint				`gorm:"check:Quantity>=0"`
	PurchaseTime	time.Time			`gorm:"index"`
}