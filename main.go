package main

import (
	"github.com/DmitryOdintsov/GinGormAPI/models"
	"github.com/DmitryOdintsov/GinGormAPI/routers"
	"github.com/DmitryOdintsov/GinGormAPI/storage"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

var err error

func main() {
	storage.DB, err = gorm.Open("postgres", "host=... user=... password=... dbname=...")
	if err != nil {
		log.Println("error DB", err)
	}
	defer func(DB *gorm.DB) {
		err := DB.Close()
		if err != nil {
			log.Println(err)
		}
	}(storage.DB)
	storage.DB.AutoMigrate(&models.Article{})

	r := routers.SetupRouter()

	//r - gin маршрутизатор
	err := r.Run()
	if err != nil {
		return
	}
}
