package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type JobApplicants struct {
	ID uuid.UUID `json:"id"`
	JobID uuid.UUID `json:"-"`
	ApplicantID uuid.UUID `json:"-"`
	Status string `json:"status"`
	Message string `json:"message"`
	Applicant *User `json:"applicant,omitempty" gorm:"foreignKey:applicant_id" `
	Job  *Job `json:"job,omitempty"`
	Audit
}


func (ja *JobApplicants) BeforeCreate(tx *gorm.DB) (err error) {
	ja.ID = uuid.New()
	return
}

func NewJobApplicants(jobID uuid.UUID, ApplicantID uuid.UUID, status string, Message string) *JobApplicants {
	return &JobApplicants{
		JobID: jobID,
		ApplicantID: ApplicantID,
		Status: status,
		Message: Message,
	}
}

func UpdateJobApplicants(id uuid.UUID, status string) *JobApplicants {
	return &JobApplicants{
		ID: id,
		Status: status,
	}
}
