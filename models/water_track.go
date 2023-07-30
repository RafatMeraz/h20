package models

import "gorm.io/gorm"

type WaterTrack struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId      uint           `json:"user_id" gorm:"index"`
	Amount      uint           `json:"amount" gorm:"default:0"`
	ConsumeTime uint           `json:"consume_time"`
	CreateAt    uint           `json:"create_at" gorm:"autoCreateTime:milli"`
	UpdatedAt   uint           `json:"updated_at" gorm:"autoUpdateTime:milli"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type WaterTrackRequest struct {
	UserId      uint
	Amount      uint `json:"amount" validate:"required"`
	ConsumeTime uint `json:"consume_time" validate:"required"`
}

type WaterTrackDTO struct {
	ID          uint `json:"id"`
	Amount      uint `json:"amount"`
	ConsumeTime uint `json:"consume_time"`
	CreateAt    uint `json:"create_at"`
	UpdatedAt   uint `json:"updated_at"`
}
