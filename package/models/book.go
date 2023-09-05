package models

import (
	"github.com/agusheryanto182/go-book-store-management-API/package/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model         // GORM.Model adalah struktur bawaan dari library GORM yang digunakan untuk mengelola bidang ID, CreatedAt, UpdatedAt, dan DeletedAt dalam tabel basis data
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&book{})
}
