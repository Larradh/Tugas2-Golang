package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

// Mahasiswa struct (Model) ...
type Mahasiswa struct {
	IDmahasiswa string        `json:"IDmahasiswa"`
	Nama        string        `json:"Nama"`
	Kelurahan   string        `json:"Kelurahan"`
	Kecamatan   string        `json:"Kecamatan"`
	Kabupaten   string        `json:"Kabupaten"`
	Provinsi    string        `json:"Provinsi"`
	Jurusan     string        `json:"Jurusan"`
	Fakultas    string        `json:"Fakultas"`
	NilaiD      []NilaiDetail `json:"NilaiD"`
}

type NilaiDetail struct {
	Matkul   string  `json:"Matkul"`
	Nilai    float64 `json:"Nilai"`
	Semester string  `json:"Semester"`
}

// Get all orders

func getMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var mhs Mahasiswa
	var nli NilaiDetail
	params := mux.Vars(r)

	sql := `SELECT
				IDmahasiswa,
				IFNULL(Nama,'') Nama,
				IFNULL(Kelurahan,'') Kelurahan,
				IFNULL(Kecamatan,'') Kecamatan,
				IFNULL(Kabupaten,'') Kabupaten,
				IFNULL(Provinsi,'')  Provinsi,
				IFNULL(Jurusan,'') Jurusan,
				IFNULL(Fakultas,'') Fakultas				
			FROM mahasiswa WHERE IDmahasiswa = ?`

	result, err := db.Query(sql, params["id"])

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&mhs.IDmahasiswa, &mhs.Nama, &mhs.Kelurahan, &mhs.Kecamatan, &mhs.Kabupaten, &mhs.Provinsi, &mhs.Jurusan, &mhs.Fakultas)

		if err != nil {
			panic(err.Error())
		}

		sqlDetial := `SELECT
						nilai.nilai
						, nilai.semester
						, matkul.matakuliah
					FROM
						nilai
						INNER JOIN matkul
						ON (nilai.idmatkul=matkul.idmatkul)
						INNER JOIN mahasiswa
						ON (nilai.idmahasiswa=mahasiswa.idmahasiswa)
					WHERE mahasiswa.idmahasiswa	= ?`

		mhsID := &mhs.IDmahasiswa
		fmt.Println(*mhsID)
		resultDetail, errDet := db.Query(sqlDetial, *mhsID)

		defer resultDetail.Close()

		if errDet != nil {
			panic(err.Error())
		}

		for resultDetail.Next() {

			err := resultDetail.Scan(&nli.Nilai, &nli.Semester, &nli.Matkul)

			if err != nil {
				panic(err.Error())
			}

			mhs.NilaiD = append(mhs.NilaiD, nli)

		}

	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mhs)

}

// Main function
func main() {

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/dbtugas")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Init router
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/mahasiswa/{id}", getMahasiswa).Methods("GET")

	fmt.Println("Server on :8080")
	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))

}
