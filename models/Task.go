package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	Id        string         `gorm:"column:id;type:uuid;primaryKey;not null;default:uuid_generate_v1()" json:"id"`
	Name      string         `gorm:"column:name" json:"name"`
	FkStatus  string         `gorm:"column:fk_status" json:"fk_status"`
	CreatedAt time.Time      `gorm:"column:created_at;default:now()" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;default:now()" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Estatus   string         `gorm:"column:estatus" json:"status"`
}
