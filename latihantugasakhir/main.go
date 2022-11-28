package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Mahasiswa
type Mahasiswa struct {
	ID   int    `json:"id"`
	NIM  int    `json:"nim"`
	Name string `json:"name"`
}

// NewMahasiswa
func NewMahasiswa() []Mahasiswa {
	mhs := []Mahasiswa{
		Mahasiswa{
			ID:   1,
			NIM:  123454,
			Name: "Didik Prabowo",
		},
		Mahasiswa{
			ID:   2,
			NIM:  923454,
			Name: "Joni Gunawan",
		},
		Mahasiswa{
			ID:   3,
			NIM:  923454,
			Name: "Muhammad Irwan",
		},
	}
	return mhs
}

// GetMahasiswa
func GetMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		mhs := NewMahasiswa()
		datamahasiswa, err := json.Marshal(mhs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(datamahasiswa)
		return
	}

	http.Error(w, "hayo mau ngapain", http.StatusNotFound)
}

// PostMahasiswa
func PostMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Mhs Mahasiswa
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			// parse dari json
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&Mhs); err != nil {
				log.Fatal(err)
			}
		} else {
			// parse dari form
			getID := r.PostFormValue("id")
			id, _ := strconv.Atoi(getID)
			getNim := r.PostFormValue("nim")
			nim, _ := strconv.Atoi(getNim)
			name := r.PostFormValue("name")
			Mhs = Mahasiswa{
				ID:   id,
				NIM:  nim,
				Name: name,
			}
		}

		dataMahasiswa, _ := json.Marshal(Mhs) // to byte
		w.Write(dataMahasiswa)                // cetak di browser
		return
	}

	http.Error(w, "hayo mau ngapain", http.StatusNotFound)
	return
}

func main() {
	http.HandleFunc("/mahasiswa", GetMahasiswa)
	// fmt.Println("server running...")
	// if err := http.ListenAndServe(":7000", nil); err != nil {
	// 	log.Fatal(err)
	// }

	http.HandleFunc("/post_mahasiswa", PostMahasiswa)
	fmt.Println("server running...")
	if err := http.ListenAndServe(":7000", nil); err != nil {
		log.Fatal(err)
	}
}