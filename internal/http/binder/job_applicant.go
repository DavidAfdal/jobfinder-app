package binder


type ApplyJobRequest struct {
	JobID string `param:"jobID"`
	Status string `json:"status"`
	Message string `json:"message"`
}

type WithdrawJobRequest struct {
	JobApplicantID string `param:"jobApplicantID"`
}

type ApproveApplicantRequest struct {
	JobApplicantID string `param:"jobApplicantID"`
}

type FindJobApplicantByIDRequest struct {
	JobApplicantID string `param:"jobApplicantID"`
}
