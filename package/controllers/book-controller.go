package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/agusheryanto182/go-book-store-management-API/package/models"
	"github.com/agusheryanto182/go-book-store-management-API/package/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// Fungsi GetBook mengambil daftar semua buku dari model dan mengirimkannya sebagai respons JSON.
func GetBook(w http.ResponseWriter, r *http.Request) {
	// Mengambil semua buku dari model
	newBooks := models.GetAllBooks()

	// Mengubah hasil menjadi format JSON
	res, _ := json.Marshal(newBooks)

	// Mengatur header respons sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Mengatur status kode OK dan mengirimkan respons
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Fungsi GetBookById mengambil detail buku berdasarkan ID yang diberikan dari model dan mengirimkannya sebagai respons JSON.
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	// Mengonversi ID buku menjadi tipe data int64
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		// Menampilkan pesan kesalahan jika terjadi kesalahan saat mengonversi
		fmt.Println("Kesalahan saat melakukan konversi")
	}

	// Mengambil detail buku berdasarkan ID dari model
	bookDetails, _ := models.GetBookById(ID)

	// Mengubah hasil menjadi format JSON
	res, _ := json.Marshal(bookDetails)

	// Mengatur header respons sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Mengatur status kode OK dan mengirimkan respons
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Fungsi CreateBook membuat buku baru berdasarkan data yang diterima dalam permintaan HTTP,
// dan kemudian mengirimkan buku yang baru dibuat sebagai respons JSON.
func CreateBook(w http.ResponseWriter, r *http.Request) {
	// Membuat instance baru dari model Book
	createBook := &models.Book{}

	// Mengurai data permintaan HTTP dan memasukkannya ke dalam instance createBook
	utils.ParseBody(r, createBook)

	// Membuat buku baru dengan memanggil fungsi CreateBook dari model
	b := createBook.CreateBook()

	// Mengubah hasil menjadi format JSON
	res, _ := json.Marshal(b)

	// Mengatur status kode OK dan mengirimkan respons
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Fungsi DeleteBook menghapus buku berdasarkan ID yang diberikan dalam permintaan HTTP,
// dan mengirimkan buku yang dihapus sebagai respons JSON.
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	// Mengonversi ID buku menjadi tipe data int64
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		// Menampilkan pesan kesalahan jika terjadi kesalahan saat mengonversi
		fmt.Println("Kesalahan saat melakukan konversi")
	}

	// Menghapus buku berdasarkan ID dari model
	book := models.DeleteBook(ID)

	// Mengubah hasil menjadi format JSON
	res, _ := json.Marshal(book)

	// Mengatur header respons sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Mengatur status kode OK dan mengirimkan respons
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Fungsi UpdateBook mengambil permintaan HTTP yang berisi data buku yang diperbarui,
// memperbarui buku dalam database berdasarkan ID yang diberikan dalam permintaan,
// dan mengirimkan buku yang telah diperbarui sebagai respons JSON.
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Membuat instance baru dari model Book untuk menyimpan data yang diperbarui
	var updateBook = &models.Book{}

	// Mengurai data permintaan HTTP dan memasukkannya ke dalam instance updateBook
	utils.ParseBody(r, updateBook)

	// Mengambil ID buku dari parameter dalam permintaan
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	// Dua angka nol ini adalah argumen tambahan yang diberikan ke fungsi ParseInt.
	//Argumen pertama adalah base (basis) untuk konversi (0 artinya akan menentukan basis dari format string),
	//dan argumen kedua adalah bit size (ukuran bit) dari tipe data yang diharapkan (0 artinya akan menghasilkan tipe data yang sesuai).
	if err != nil {
		// Menampilkan pesan kesalahan jika terjadi kesalahan saat mengonversi
		fmt.Println("Kesalahan saat melakukan konversi")
	}

	// Mengambil detail buku berdasarkan ID dari model
	booksDetails, db := models.GetBookById(ID)

	// Memeriksa apakah ada data yang diperbarui untuk buku
	if updateBook.Name != "" {
		booksDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		booksDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		booksDetails.Publication = updateBook.Publication
	}

	// Menyimpan perubahan ke dalam database
	db.Save(&booksDetails)

	// Mengubah hasil menjadi format JSON
	res, err := json.Marshal(booksDetails)
	if err != nil {
		fmt.Println("error dibagian konversi ke json")
	}

	// Mengatur header respons sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Mengatur status kode OK dan mengirimkan respons
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
