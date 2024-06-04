package service

import (
	"errors"

	"github.com/DavidAfdal/workfinder/internal/entity"
	"github.com/DavidAfdal/workfinder/internal/repository"
	"github.com/google/uuid"
)

type JobApplicantService interface {
	ApplyJob(jobApplicant *entity.JobApplicants) (*entity.JobApplicants, error)
	WithdrawJob(id uuid.UUID, userID uuid.UUID) (bool, error)
	ApproveApplicant(id uuid.UUID, userID uuid.UUID) (*entity.JobApplicants, error)
	FindJobApplicantByID(id uuid.UUID) (*entity.JobApplicants, error)
}

type jobApplicantService struct {
	jobApplicantRepo repository.JobApplicantsRepository
	jobRepo          repository.JobRepository
}

func NewJobApplicantService(jobApplicantRepo repository.JobApplicantsRepository, jobRepo repository.JobRepository) JobApplicantService {
	return &jobApplicantService{jobApplicantRepo,jobRepo}
}

func (s *jobApplicantService) ApplyJob(jobApplicant *entity.JobApplicants) (*entity.JobApplicants, error) {

	job, err := s.jobRepo.FindJobByID(jobApplicant.JobID)

	if err != nil {
		return nil, errors.New("job not found")
	}

	if job.Closed == true {
		return nil, errors.New("job already closed")
	}

	if jobApplicant.ApplicantID == job.ClientID {
		return nil, errors.New("you can't apply for your own job")
	}

	jobApplicantData, _ := s.jobApplicantRepo.FindJobApplicantsByID(jobApplicant.ID)

	if jobApplicantData != nil {
		return jobApplicantData, errors.New("job already applied")
	}

	return s.jobApplicantRepo.ApplyJob(jobApplicant)
}

func (s *jobApplicantService) WithdrawJob(id uuid.UUID, userID uuid.UUID) (bool, error) {
	jobApplicant, err := s.jobApplicantRepo.FindJobApplicantsByID(id)

	if err != nil {
		return false, err
	}

	if jobApplicant.ApplicantID != userID {
		return false, errors.New("unauthorized")
	}

	return s.jobApplicantRepo.WithdrawJob(jobApplicant)
}

func (s *jobApplicantService) ApproveApplicant(id uuid.UUID, userID uuid.UUID) (*entity.JobApplicants, error) {


	jobApplicant, err := s.jobApplicantRepo.FindJobApplicantsByID(id)

	job, err := s.jobRepo.FindJobByID(jobApplicant.JobID)

	if err != nil {
		return jobApplicant, err
	}


	if job.Closed == true {
		return jobApplicant, errors.New("job already closed")
	}


	if job.Client.ID != userID {
		return jobApplicant, errors.New("unauthorized")
	}

	if jobApplicant.ApplicantID == userID {
		return jobApplicant, errors.New("can't approve yourself")
	}

	return s.jobApplicantRepo.ApproveApplicant(jobApplicant)
}


func (s *jobApplicantService) FindJobApplicantByID(id uuid.UUID) (*entity.JobApplicants, error) {
	return s.jobApplicantRepo.FindJobApplicantsByID(id)
}
