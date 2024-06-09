package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)



type User struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	Email string `json:"email,omitempty"`
	Password string `json:"-"`
	Address string `json:"address,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Gender string `json:"gender,omitempty"`
	Role string `json:"-"`
	Audit
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

func NewUser(name string, email string, password string, address string, phoneNumber string, gender string, Role string) *User {
	return &User{
		Name: name,
		Email: email,
		Password: password,
		Address: address,
		PhoneNumber: phoneNumber,
		Gender: gender,
		Role: Role,
		Audit: NewAuditTable(),
	}
}

func UpdateUser(id uuid.UUID,name string, email string, password string, address string, phoneNumber string, gender string ) *User {
	return &User{
		ID: id,
		Name: name,
		Email: email,
		Password: password,
		Address: address,
		PhoneNumber: phoneNumber,
		Gender: gender,
		Audit: UpdateAuditTable(),
	}
}
