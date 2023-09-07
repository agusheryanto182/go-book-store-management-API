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

// Inisialisasi database dan migrasi otomatis tabel Book saat aplikasi dimulai
func init() {
	config.Connect()        // Menghubungkan ke database
	db = config.GetDB()     // Mendapatkan instance database
	db.AutoMigrate(&Book{}) // Melakukan migrasi otomatis tabel Book
}

// Membuat buku baru dan menyimpannya ke dalam database
func (b *Book) CreateBook() *Book {
	db.NewRecord(b) // Menandai buku sebagai data yang baru
	db.Create(b)    // Menyimpan buku ke dalam database
	return b
}

// Mengambil daftar semua buku dari database
func GetAllBooks() []Book {
	var books []Book
	db.Find(&books) // Mengambil semua buku dari database
	return books
}

// Mengambil buku berdasarkan ID yang diberikan
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID = ?", Id).Find(&getBook) // Mencari buku berdasarkan ID
	return &getBook, db
}

// Menghapus buku berdasarkan ID yang diberikan
func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID = ?", ID).Delete(book) // Menghapus buku dari database
	return book
}
