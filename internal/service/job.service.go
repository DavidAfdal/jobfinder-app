package service

import (
	"github.com/DavidAfdal/workfinder/internal/entity"
	"github.com/DavidAfdal/workfinder/internal/repository"
	"github.com/google/uuid"
)

// TODO: Create Job Service Struct and Interface

// TODO: Create Job Service Implementation


type JobService interface {
	FindAllJob() ([]entity.Job, error)
	FindJobByID(id uuid.UUID) (*entity.Job, error)
	FindSharedJobs(userID uuid.UUID) ([]entity.Job, error)
	FindAppliedJobs(userID uuid.UUID) ([]entity.Job, error)
	CreateJob(job *entity.Job) (*entity.Job, error)
	UpdateJob(job *entity.Job) (*entity.Job, error)
	DeleteJob(id uuid.UUID) (bool, error)
}

type jobService struct {
	jobRepo repository.JobRepository
}


func NewJobService(jobRepo repository.JobRepository) JobService {
	return &jobService{jobRepo: jobRepo}
}


func (s *jobService) FindAllJob() ([]entity.Job, error) {
	return s.jobRepo.FindAllJob()
}


func (s *jobService) FindJobByID(id uuid.UUID) (*entity.Job, error) {
	return s.jobRepo.FindJobByID(id)
}

func (s *jobService) FindSharedJobs(userID uuid.UUID) ([]entity.Job, error) {
	return s.jobRepo.FindSharedJob(userID)
}
func (s *jobService) FindAppliedJobs(userID uuid.UUID) ([]entity.Job, error) {
	return s.jobRepo.FindAppliedJob(userID)
}
func (s *jobService) CreateJob(job *entity.Job) (*entity.Job, error) {
	return s.jobRepo.CreateJob(job)
}

func (s *jobService) UpdateJob(job *entity.Job) (*entity.Job, error) {
	return s.jobRepo.UpdateJob(job)
}

func (s *jobService) DeleteJob(id uuid.UUID)  (bool, error) {
	job, err := s.jobRepo.FindJobByID(id)

	if err != nil {
		return false, err
	}

	return s.jobRepo.DeleteJob(job)
}
