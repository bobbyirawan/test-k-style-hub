package entity

import (
	"time"

	"gorm.io/gorm"
)

type Entity struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
