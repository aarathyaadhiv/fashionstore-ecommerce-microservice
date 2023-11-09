package db

import (
	"log"

	"github.com/aarathyaadhiv/fashionstore-product-service/pkg/config"
	"github.com/aarathyaadhiv/fashionstore-product-service/pkg/domain"
	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)



func ConnectDatabase(c config.Config)(*gorm.DB,error){
	db,err:=gorm.Open(postgres.Open(c.DBUrl),&gorm.Config{})

	if err!=nil{
		log.Fatalln(err)
	}

	
	db.AutoMigrate(&domain.Products{})
	
	db.AutoMigrate(&domain.Category{})
	db.AutoMigrate(&domain.Brand{})

	return db,err
}