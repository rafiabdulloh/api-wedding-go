package models

import (
	"gorm.io/gorm"
	"time"

)

type Informasi_mempelai struct {
	gorm.Model
	// Id int64  `gorm:"primaryKey" json:"id"`
	Namapria string `gorm:"type:varchar(300)" json:"namapria"`
	Namawanita string `gorm:"type:varchar(300)" json:"namawanita"`
	Namaibupria string `gorm:"type:varchar(300)" json:"namaibupria"`
	Namabapakpria string `gorm:"type:varchar(300)" json:"namabapakpria"`
	Namaibuwanita string `gorm:"type:varchar(300)" json:"namaibuwanita"`
	Namabapakwanita string `gorm:"type:varchar(300)" json:"namabapakwanita"`
	Kode int `gorm:"type:integer" json:"kode"`
	Alamat string `gorm:"type:text" json:"alamat"`
	Waktuacara time.Time `json:"waktuacara"`
	Pukulakad string `gorm:"type:varchar(300)" json:"pukulakad"`
	Pukulresepsi string `gorm:"type:varchar(300)" json:"pukulresepsi"`
}

type Nomor_rekening struct {
	gorm.Model
	// Id int64 `gorm:"primaryKey" json:"id"`
	Namawallet string `gorm:"type:varchar(300)" json:"namawallet"`
	Nomorwallet int64 `gorm:"type:integer" json:"nomerwallet"`
	Kode int `gorm:"type:integer" json:"kode"`
}

type Komentar struct {
	gorm.Model
	// Id int64 `gorm:"primaryKey" json:"id"`
	Nama string `gorm:"type:varchar(300)" json:"nama"`
	Ucapan string `gorm:"type:text" json:"ucapan"`
	Kehadiran bool `json:"kehadiran"`
	Kode int `gorm:"type:integer" json:"kode"`
}

type Gallery struct {
	gorm.Model
	// Id int64 `gorm:"primaryKey" json:"id"`
	Label string `gorm:"type:varchar(300)" json:"label"`
	File string `gorm:"type:varchar(300)" json:"file"`
	Kode int `gorm:"type:integer" json:"kode"`
}