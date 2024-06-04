package binder

import "github.com/google/uuid"

type JobFindByIDRequest struct {
	ID string `param:"id" validate:"required"`
}
type FindSharedJobsRequest struct {
	UserID string `param:"userID" validate:"required"`
}


type CreateJobRequest struct {
	Title 		string `json:"title"`
	Description string `json:"description"`
	Company 	string `json:"company"`
	Logo 		string `json:"logo"`
	Status 		string `json:"status"`
	Salary 		float64    `json:"salary"`
	Location 	string `json:"location"`
	CategoryID  uuid.UUID `json:"category_id"`
}

type UpdateJobRequest struct {
	ID 			uuid.UUID `param:"id"`
	Title 		string `json:"title"`
	Description string `json:"description"`
	Company 	string `json:"company"`
	Logo 		string `json:"logo"`
	Status 		string `json:"status"`
	Salary 		float64    `json:"salary"`
	Location 	string `json:"location"`
	CategoryID  uuid.UUID `json:"category_id"`
	ClientID 	uuid.UUID `json:"client_id"`
}
