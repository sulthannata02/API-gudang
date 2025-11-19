package models

type Barang struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Nama     string `json:"nama"`
	Stok     int    `json:"stok"`
	Lokasi   string `json:"lokasi"`
}
