package model

type Product struct {
	ID 			string				`gorm:"primaryKey"`
	Name 	  	string				`gorm:"unique"`
	Description string
	Price 		uint
	Quantity 	uint				`gorm:"check:quantity>=0"`
}
