package main

import (
	"github.com/agusheryanto182/go-book-store-management-API/package/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Membuat router menggunakan mux.NewRouter() dari paket "mux".
	r := mux.NewRouter()

	// Mendaftarkan rute-rute untuk toko buku menggunakan fungsi RegisterBookStoreRoutes
	// dari paket "routes". Ini akan menghubungkan URL dengan handler yang sesuai.
	routes.RegisterBookStoreRoutes(r)

	// Mengatur router sebagai penanganan utama untuk semua permintaan web.
	http.Handle("/", r)

	// Memulai server web dan mendengarkan di alamat "localhost:8080".
	// Jika ada kesalahan dalam menjalankan server, log.Fatal akan digunakan
	// untuk mengakhiri program dan mencetak pesan kesalahan.
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
