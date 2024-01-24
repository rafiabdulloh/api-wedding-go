package models

import (
	"fmt"
	"log"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	// "database/sql"
	"gorm.io/driver/postgres"
	"time"

)

var DB *gorm.DB

func GetConnect() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "postgres"
		dbname   = "pernikahan"
	)

	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatal("err env")
		return
	}

	// Buat string koneksi PostgreSQL
    dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }
    // defer db.Close()

	db.AutoMigrate(&Informasi_mempelai{})
	db.AutoMigrate(&Nomor_rekening{})
	db.AutoMigrate(&Komentar{})
	db.AutoMigrate(&Gallery{})

	DB = db

}

func Seeder() {

	waktuAcara := time.Date(2027, time.February, 6, 0, 0, 0, 0, time.UTC)
	dataInformasi_mempelai := Informasi_mempelai{
		Namapria : "Sukma",
		Namawanita : "Sukmi",
		Namaibupria : "Mom Of Sukma",
		Namabapakpria : "Dad Of Sukma",
		Namaibuwanita : "Mom Of Sukmi",
		Namabapakwanita : "Dad Of Sukmi",
		Waktuacara : waktuAcara,
		Pukulakad : "10:00",
		Pukulresepsi: "08:00",
		Kode :123456,
		Alamat: "Kab. Bandung Barat, Kec. Cikalong Wetan, Desa Ciptagumati, Kp. Cinangsi RT 01/08",
	}

	dataNomor_rekening:= Nomor_rekening {
		Namawallet : "Wahyu",
		Nomorwallet : 0011223344,
		Kode : 123456,
	}

	dataKomentar := Komentar {
		Nama : "rafi",
		Ucapan : "Selamat atas Pernikahannya",
		Kehadiran : true, 
		Kode : 123456,
	}
	
	dataGallery := Gallery {
		Label : "fotbar kota baru",
		File : "image/0099988779.jpg",
		Kode : 123456,
	}

	entitiesToCreate := []interface{}{&dataNomor_rekening, &dataKomentar, &dataGallery, &dataInformasi_mempelai}
	for _, entity := range entitiesToCreate {
		var count int64
		DB.Model(&Informasi_mempelai{}).Count(&count)
		if count == 0 {
			// Jika database kosong, maka insert dataInformasi_mempelai awal
			result := DB.Create(entity)
			if result.Error != nil {
				fmt.Printf("Error creating entity: %v\n", result.Error)
			}
			fmt.Println("========seeder running=============")
		}
		
	}

}
