package models

import "gorm.io/gorm"

type Transaksi struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	IDBarang   uint   `json:"id_barang"`
	Jenis      string `json:"jenis"` // masuk / keluar
	Jumlah     int    `json:"jumlah"`
	Keterangan string `json:"keterangan"`
	gorm.Model
}
