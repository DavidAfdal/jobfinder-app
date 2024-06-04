package repository

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/DavidAfdal/workfinder/internal/entity"
	"github.com/DavidAfdal/workfinder/pkg/cache"
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type JobRepository interface {
	FindAllJob() ([]entity.Job, error)
	FindJobByID(id uuid.UUID) (*entity.Job, error)
	FindSharedJob(userId uuid.UUID) ([]entity.Job, error)
	FindAppliedJob(userId uuid.UUID) ([]entity.Job, error)
	CreateJob(job *entity.Job) (*entity.Job, error)
	UpdateJob(job *entity.Job) (*entity.Job, error)
	DeleteJob(job *entity.Job) (bool, error)
}

type jobRepository struct {
	db *gorm.DB
	cahce cache.Cacheable
}


func NewJobRepository(db *gorm.DB, cahce cache.Cacheable) JobRepository {
	return &jobRepository{db, cahce}
}

func (r *jobRepository) FindAllJob() ([]entity.Job, error) {
	// TODO: implement find all Jobs method
	jobs := make([]entity.Job, 0)

	key:= "GetAllJobs"

	data := r.cahce.Get(key)

	if data == "" {
		if err := r.db.Preload("Category", func (db *gorm.DB) *gorm.DB {
			return db.Select("title", "id", "icon")
		}).Find(&jobs).Error; err != nil {
			return jobs, err
		}

		marshalJob, _:= json.Marshal(jobs)
		err := r.cahce.Set(key, marshalJob, 2 * time.Minute)

		if err != nil {
			return jobs, err
		}
	} else {
		err := json.Unmarshal([]byte(data), &jobs)
		if err != nil {
			return jobs, err
		}
	}


	return jobs, nil
}

func (r *jobRepository) FindSharedJob(userId uuid.UUID) ([]entity.Job, error) {
	jobs := make([]entity.Job, 0)

	key:= "GetSharedJobs"

	data := r.cahce.Get(key)

	if data == "" {
		if err := r.db.Preload("Category", func (db *gorm.DB) *gorm.DB {
			return db.Select("title", "id", "icon")
		}).Find(&jobs, "client_id = ?", userId).Error; err != nil {
			return jobs, err
		}

		marshalJob, _:= json.Marshal(jobs)
		err := r.cahce.Set(key, marshalJob, 2 * time.Minute)

		if err != nil {
			return jobs, err
		}
	} else {
		err := json.Unmarshal([]byte(data), &jobs)
		if err != nil {
			return jobs, err
		}
	}


	return jobs, nil

}

func (r *jobRepository) FindJobByID(id uuid.UUID) (*entity.Job, error) {
	job := new(entity.Job)


	key := fmt.Sprintf("job_%s", id)

	data := r.cahce.Get(key)

	if data == "" {
		if err := r.db.Preload("Applicants", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Applicant", func(db *gorm.DB) *gorm.DB {
				return db.Select("name", "id")
			})
   		}).
		Preload("Category", func (db *gorm.DB) *gorm.DB {
			return db.Select("title", "id", "icon")
		}).
		Preload("Client", func(db *gorm.DB) *gorm.DB {
			return db.Select("id","name", "email")
		}).
		Where("id = ?", id).
		First(&job).Error;
		err != nil {
			return job, err
		}

		if len(job.Applicants) == 0 {
			job.Applicants = make([]*entity.JobApplicants, 0)
		}

		marshalJob, _:= json.Marshal(job)
		err := r.cahce.Set(key, marshalJob, 2 * time.Minute)

		if err != nil {
			return job, err
		}
	} else {
		err := json.Unmarshal([]byte(data), &job)
		if err != nil {
			return job, err
		}
	}

	return job, nil
}

func (r *jobRepository) FindAppliedJob(userId uuid.UUID) ([]entity.Job, error){

	jobs := make([]entity.Job, 0)
	applicant_jobs := make([]entity.JobApplicants, 0)

	if err := r.db.Where("applicant_id = ?", userId).Preload("Job.Category").Find(&applicant_jobs).Error; err != nil {
		return jobs, err
	}


	for _, v := range applicant_jobs {
		jobs = append(jobs, *v.Job)
	}

	return jobs, nil
}


func (r *jobRepository) CreateJob(job *entity.Job) (*entity.Job, error) {
	if err := r.db.Create(&job).Error; err != nil {
		return job, err
	}

	return job, nil
}

func (r *jobRepository) UpdateJob(job *entity.Job) (*entity.Job, error) {
	fields := make(map[string]interface{})

	if job.Title != "" {
		fields["title"] = job.Title
	}

	if job.Description != "" {
		fields["description"] = job.Description
	}

	if job.Logo != "" {
		fields["logo"] = job.Logo
	}

	if job.Status != "" {
		fields["status"] = job.Status
	}

	if job.Location != "" {
		fields["location"] = job.Location
	}

	if job.Company != "" {
		fields["company"] = job.Company
	}

	if job.Salary != 0 {
		fields["salary"] = job.Salary
	}

	if err := r.db.Model(&job).Updates(fields).Error; err != nil {
		return job, err
	}

	return job, nil
}


func (r *jobRepository) DeleteJob(job *entity.Job) (bool, error){
	if err:= r.db.Delete(&job).Error; err != nil {
		return false, nil
	}
	return true, nil
}
