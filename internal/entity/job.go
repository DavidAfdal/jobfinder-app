package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type Job struct {
	ID 			uuid.UUID `json:"id"`
	Title 		string `json:"title"`
	Description string `json:"description,omitempty"`
	Company 	string `json:"company,omitempty"`
	Logo 		string `json:"logo,omitempty"`
	Status 		string `json:"status,omitempty"`
	Salary 		float64 `json:"salary,omitempty"`
	Location 	string `json:"location,omitempty"`
	Closed    	bool   `json:"closed,omitempty"`
	CategoryID  uuid.UUID `json:"-"`
	ClientID 	uuid.UUID `json:"-"`
	Category    *Category `json:"category,omitempty"`
	Client      *User     `json:"client,omitempty" gorm:"foreignKey:client_id"`
	Applicants []*JobApplicants `json:"applicants,omitempty"`
	Audit
}

func (j *Job) BeforeCreate(tx *gorm.DB) (err error) {
	j.ID = uuid.New()
	return
}


func NewJob(title string, description string, company string, logo string, status string, salary float64, location string, categoryID uuid.UUID, clientID uuid.UUID) *Job {
	return &Job{
		Title: title,
		Description: description,
		Company: company,
		Logo: logo,
		Status: status,
		Salary: salary,
		Closed: false,
		Location: location,
		CategoryID: categoryID,
		ClientID: clientID,
		Audit: NewAuditTable(),
	}
}

func UpdateJob(id uuid.UUID, title string, description string, company string, logo string, status string, salary float64, location string) *Job {
	return &Job{
		ID: id,
		Title: title,
		Description: description,
		Company: company,
		Logo: logo,
		Status: status,
		Salary: salary,
		Location: location,
		Audit: UpdateAuditTable(),
	}
}
