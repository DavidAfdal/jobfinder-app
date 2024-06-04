package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type Category struct {
	ID    uuid.UUID `json:"id"`
	Title string `json:"title"`
	Icon string `json:"icon"`
	Jobs []*Job	`json:"jobs,omitempty"`
	Audit
}


func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}


func NewCategory(title string, icon string) *Category {
	return &Category{
		Title: title,
		Icon: icon,
		Audit: NewAuditTable(),
	}
}

func UpdateCategory(id uuid.UUID, title string, icon string) *Category {
	return &Category{
		ID: id,
		Title: title,
		Icon: icon,
		Audit: UpdateAuditTable(),
	}
}
