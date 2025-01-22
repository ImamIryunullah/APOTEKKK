package models

import (
    "time"
)

type Stok struct {
    ID             uint      `json:"id" gorm:"primaryKey"`
    ObatID         uint      `json:"obat_id" gorm:"not null"`
    Jumlah         int       `json:"jumlah" gorm:"not null"`
    TipeTransaksi  string    `json:"tipe_transaksi" gorm:"type:enum('MASUK', 'KELUAR');not null"` 
    Keterangan     string    `json:"keterangan" gorm:"type:text"`
    CreatedAt      time.Time `json:"created_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
    UpdatedAt      time.Time `json:"updated_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP;autoUpdateTime"`

    Obat           Obat      `json:"obat" gorm:"foreignKey:ObatID;references:ID"`
}
