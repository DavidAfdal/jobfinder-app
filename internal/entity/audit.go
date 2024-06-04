package entity

import (
	"time"

	"gorm.io/gorm"
)

type Audit struct {
	CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"-"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}


func NewAuditTable() Audit{
	return Audit{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func UpdateAuditTable() Audit{
	return Audit{
		UpdatedAt: time.Now(),
	}
}
