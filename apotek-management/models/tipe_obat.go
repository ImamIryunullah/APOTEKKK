// models/tipe_obat.go

package models

import "time"

type TipeObat struct {
    ID        uint      `json:"id" gorm:"primaryKey;type:bigint unsigned"`
    NamaTipe  string    `json:"nama_tipe" gorm:"type:varchar(100);not null"`
    CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
    UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP;autoUpdateTime"`
} 

func (TipeObat) TableName() string {
    return "tipe_obats"
}