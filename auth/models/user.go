package models

import "gorm.io/gorm"

type User struct {
	ID        int            `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt uint           `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdateAt  uint           `json:"update_at" gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
