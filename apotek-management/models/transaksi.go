package models

import (
	"time"
)

type Transaksi struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	KodeTransaksi string    `json:"kode_transaksi" gorm:"type:varchar(20);unique;not null"`
	Jumlah        int64     `json:"jumlah" gorm:"type:bigint;not null"`
	HargaSatuan   float64   `json:"harga_satuan" gorm:"type:decimal(10,2);not null"`
	TotalHarga    float64   `json:"total_harga" gorm:"type:decimal(10,2);not null"`
	Status        string    `json:"status" gorm:"type:varchar(50);not null"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"` 
}


func (transaksi *Transaksi) BeforeCreate() {
	// Hitung total harga saat transaksi dibuat
	transaksi.TotalHarga = float64(transaksi.Jumlah) * transaksi.HargaSatuan
}
