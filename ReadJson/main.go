package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

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

// NilaiDetail struct (Model) ...
type NilaiDetail struct {
	Matkul   string  `json:"Matkul"`
	Nilai    float64 `json:"Nilai"`
	Semester string  `json:"Semester"`
}

func main() {

	url := "http://localhost:8080/mahasiswa"

	spaceClient := http.Client{
		Timeout: time.Second * 2, //Timeout after 2 second

	}
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		log.Fatal(readErr)
	}

	mhs := Mahasiswa{}
	jsonErr := json.Unmarshal(body, &mhs)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(mhs.IDmahasiswa)
	fmt.Println(mhs.Nama)

	for _, nilai := range mhs.NilaiD {
		fmt.Println("NamaMatkul", nilai.Nilai)
		fmt.Println("Nilai", nilai.Matkul)
		fmt.Println("Semester", nilai.Semester)

	}

}
