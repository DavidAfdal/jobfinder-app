package repository

import (
	"github.com/DavidAfdal/workfinder/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type JobApplicantsRepository interface {
	ApplyJob(jobApplicant *entity.JobApplicants) (*entity.JobApplicants, error)
	FindJobApplicantsByID(id uuid.UUID) (*entity.JobApplicants, error)
	WithdrawJob(jobApplicant *entity.JobApplicants) (bool, error)
	ApproveApplicant(jobApplicants *entity.JobApplicants) (*entity.JobApplicants, error)
}

type jobApplicantsRepository struct {
	db *gorm.DB
}


func NewJobApplicantsRepository(db *gorm.DB) JobApplicantsRepository {
	return &jobApplicantsRepository{db}
}

func (r *jobApplicantsRepository) ApplyJob(jobApplicant *entity.JobApplicants) (*entity.JobApplicants, error) {
	if err := r.db.Create(&jobApplicant).Error; err != nil {
		return jobApplicant, err
	}

	return jobApplicant, nil
}

func (r *jobApplicantsRepository) FindJobApplicantsByID(id uuid.UUID) (*entity.JobApplicants, error)  {
	jobApplicant := new(entity.JobApplicants)
	if err := r.db.Preload("Applicant", func(db *gorm.DB) *gorm.DB{
		return db.Select("name", "id", "email")
	}).First(&jobApplicant, "id = ?", id).Error; err != nil {
		return jobApplicant, err
	}

	return jobApplicant, nil
}

func (r *jobApplicantsRepository) WithdrawJob(jobApplicant *entity.JobApplicants) (bool, error)  {
	if err := r.db.Delete(&jobApplicant).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *jobApplicantsRepository) ApproveApplicant(jobApplicants *entity.JobApplicants) (*entity.JobApplicants, error) {

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := r.db.Model(&jobApplicants).Update("status", "Approved").Error; err != nil {
			return err
		}

		if err := r.db.Model(&entity.JobApplicants{}).Where("id != ?", jobApplicants.ID).Update("status", "Rejected").Error; err != nil {
			return err
		}

		if err := r.db.Model(&entity.Job{}).Where("id = ?", jobApplicants.Job.ID).Update("closed", true).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return jobApplicants, err
	}

	return jobApplicants, nil
}
