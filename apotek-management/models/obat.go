package models

import "time"

type Obat struct {
    ID         uint      `json:"id" gorm:"primaryKey"`
    KodeObat   string    `json:"kode_obat"`
    NamaObat   string    `json:"nama_obat" gorm:"type:varchar(100);not null"`
    Deskripsi  string    `json:"deskripsi" gorm:"type:text"`
    Harga      float64   `json:"harga" gorm:"type:decimal(10,2);not null"`
    StokAwal   int       `json:"stok_awal" gorm:"type:int;not null"`
    TipeObatID uint      `json:"tipe_obat_id" gorm:"type:bigint unsigned;not null"`
    CreatedAt  time.Time `json:"created_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
    UpdatedAt  time.Time `json:"updated_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP;autoUpdateTime"`

    TipeObat TipeObat `json:"tipe_obat" gorm:"foreignKey:TipeObatID;references:ID"`
}
