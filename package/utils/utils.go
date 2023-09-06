package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	// Membaca seluruh isi body dari permintaan HTTP
	body, err := ioutil.ReadAll(r.Body)

	// Memeriksa apakah pembacaan body berhasil
	if err == nil {
		// Menguraikan isi body yang berisi JSON ke dalam variabel x yang ditentukan oleh pengguna
		if err := json.Unmarshal([]byte(body), x); err != nil {
			// Jika terjadi kesalahan dalam penguraian JSON, fungsi ini mengembalikan nilai
			return
		}
	}
}
