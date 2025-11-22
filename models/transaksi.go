package models

import "gorm.io/gorm"

type Transaksi struct {
	gorm.Model
	IDBarang   uint   `json:"id_barang"`
	Jenis      string `json:"jenis"`
	Jumlah     int    `json:"jumlah"`
	Keterangan string `json:"keterangan"`
}
