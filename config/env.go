package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("WARNING: .env tidak ditemukan atau gagal dibaca")
	}
}
