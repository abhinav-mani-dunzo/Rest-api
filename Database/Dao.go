package Database

import (
	"Rest-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var dao *gorm.DB

func Get() (*gorm.DB, error) {
	log.Println("Dao get called")
	if dao != nil {
		return dao, nil
	}
	dsn := "host=localhost user=abhinav password=abhi1946 dbname=shop port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err!=nil{
		log.Fatal(err)
		return nil, err
	}else{
		db.AutoMigrate(&model.Product{})
		db.AutoMigrate(&model.Purchase{})
	}
	dao = db
	return dao , nil
}
