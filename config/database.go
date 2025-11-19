package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gudang-app/models"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:@tcp(127.0.0.1:3306)/gudang_db?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}

	DB = db
}

func AutoMigrate() {
	DB.AutoMigrate(
		&models.User{},
		&models.Barang{},
		&models.Transaksi{},
	)
}
