package handler

import (
	"net/http"

	"github.com/DavidAfdal/workfinder/internal/entity"
	"github.com/DavidAfdal/workfinder/internal/http/binder"
	"github.com/DavidAfdal/workfinder/internal/service"
	"github.com/DavidAfdal/workfinder/pkg/response"
	"github.com/DavidAfdal/workfinder/pkg/token"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)


type JobApplicantsHandler interface {
	ApplyJob(ctx echo.Context) error
	WithdrawnJobApplicants(ctx echo.Context) error
	ApproveApplicant(ctx echo.Context) error
	FindJobApplicantsByID(ctx echo.Context) error
}

type jobApplicantsHandler struct {
	jobApplicantsService service.JobApplicantService
}

func NewJobApplicantsHandler(jobApplicantsService service.JobApplicantService) JobApplicantsHandler {
	return &jobApplicantsHandler{jobApplicantsService}
}


// ApplyJob godoc
// @Summary Apply for a job
// @Description Apply for a job
// @Tags job-applicants
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param applyJob body binder.ApplyJobRequest true "Apply Job Request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /job-applicants/apply [post]
func (h *jobApplicantsHandler) ApplyJob(ctx echo.Context) error {

	dataUser, _ := ctx.Get("user").(*jwt.Token)
	claims := dataUser.Claims.(*token.JwtCustomClaims)

	var input binder.ApplyJobRequest

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	jobID := uuid.MustParse(input.JobID)

	newJobApplicant := entity.NewJobApplicants(jobID, claims.ID, input.Status, input.Message)


	_, err := h.jobApplicantsService.ApplyJob(newJobApplicant)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}


	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success apply job", nil))
}


// WithdrawnJobApplicants godoc
// @Summary Withdraw job application
// @Description Withdraw job application
// @Tags job-applicants
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param withdrawJob body binder.WithdrawJobRequest true "Withdraw Job Request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /job-applicants/withdraw [post]
func (h *jobApplicantsHandler) WithdrawnJobApplicants(ctx echo.Context) error {
	dataUser, _ := ctx.Get("user").(*jwt.Token)
	claims := dataUser.Claims.(*token.JwtCustomClaims)

	var input binder.WithdrawJobRequest

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}


	id := uuid.MustParse(input.JobApplicantID)
	_, err := h.jobApplicantsService.WithdrawJob(id, claims.ID)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success withdraw job", nil))
}

// FindJobApplicantsByID godoc
// @Summary Find job applicant by ID
// @Description Find job applicant by ID
// @Tags job-applicants
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "Job Applicant ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /job-applicants/{id} [get]
func (h *jobApplicantsHandler) FindJobApplicantsByID(ctx echo.Context) error {

	var input binder.FindJobApplicantByIDRequest

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}


	id := uuid.MustParse(input.JobApplicantID)

	jobApplicant, err := h.jobApplicantsService.FindJobApplicantByID(id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success find job applicant", jobApplicant))

}

// ApproveApplicant godoc
// @Summary Approve job applicant
// @Description Approve job applicant
// @Tags job-applicants
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param approveApplicant body binder.ApproveApplicantRequest true "Approve Applicant Request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /job-applicants/approve [post]
func (h *jobApplicantsHandler) ApproveApplicant(ctx echo.Context) error {
	dataUser, _ := ctx.Get("user").(*jwt.Token)
	claims := dataUser.Claims.(*token.JwtCustomClaims)

	var input binder.ApproveApplicantRequest

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	id := uuid.MustParse(input.JobApplicantID)


	_, err := h.jobApplicantsService.ApproveApplicant(id, claims.ID)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success approve job", nil))
}

