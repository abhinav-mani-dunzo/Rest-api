package Service

import (
	"Rest-api/Database"
	"Rest-api/model"
	"gorm.io/gorm"
	"log"
	"time"
)

func GetAllProduct() []model.Product {
	var products []model.Product
	doa, err := Database.Get()
	if err != nil {
		return nil
	}
	doa.Find(&products)
	return products
}

func GetProductById(id string) (model.Product, error) {
	log.Printf("Searching product by id: %s\n", id)
	var product model.Product
	doa, err := Database.Get()
	if err != nil {
		log.Println(err)
	}
	err = doa.First(&product, id).Error
	return product, err
}

func AddAllProducts(products []model.Product) error {
	log.Printf("Updating Following Product\n %s \n", products)
	doa, err := Database.Get()
	if err != nil {
		log.Println("Error", err)
	}
	err = doa.Save(&products).Error
	return err
}

func Purchase(purchase model.Purchase) error {
	dao, err := Database.Get()
	if err != nil {
		log.Fatalln(err)
	}
	err = dao.Model(&model.Product{}).Where("id = ?", purchase.Item).Update("quantity", gorm.Expr("quantity - ?", purchase.Quantity)).Error
	if err != nil {
		log.Println(err)
		return err
	}
	purchase.PurchaseTime = time.Now()
	err = dao.Create(&purchase).Error
	return err
}

func TopFiveProduct() []model.Product {
	var topfive []model.Product
	var topfiveid []string
	dao, err := Database.Get()
	if err != nil {
		log.Fatalln(err)
	}
	dao.Model(&model.Purchase{}).Select("item").Where("purchase_time > ?", time.Now().Add(time.Hour*-1)).Group("item").Order("sum(quantity) desc").Limit(5).Find(&topfiveid)
	dao.Model(&model.Product{}).Where("id", topfiveid).Find(&topfive)
	return topfive
}
