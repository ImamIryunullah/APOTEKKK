package models

import "time"

type TagObat struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	NamaTag   string    `json:"nama_tag" gorm:"type:varchar(50);not null;unique"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime"`
}

func (TagObat) TableName() string {
	return "tag_obat"
}
