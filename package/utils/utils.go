package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	// Membaca seluruh isi body dari permintaan HTTP
	body, err := io.ReadAll(r.Body)

	// Memeriksa apakah pembacaan body berhasil
	if err == nil {
		// Menguraikan isi body yang berisi JSON ke dalam variabel x yang ditentukan oleh pengguna
		if err := json.Unmarshal(body, x); err != nil {
			// Jika terjadi kesalahan dalam penguraian JSON, fungsi ini mengembalikan nilai
			fmt.Println("terjadi kesalahan")
			return
		}
	}
}
