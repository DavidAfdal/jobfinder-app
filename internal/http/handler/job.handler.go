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


type JobHandler interface {
    FindJobs(ctx echo.Context) error
	FindJobByID(ctx echo.Context) error
	CreateJob(ctx echo.Context) error
	UpdateJob(ctx echo.Context) error
	FindSharedJobs(ctx echo.Context)error
	FindAppliedJobs(ctx echo.Context)error
	DeleteJob(ctx echo.Context) error
}

type jobHandler struct {
	jobService service.JobService
}


func NewJobHandler(jobService service.JobService) JobHandler {
	return &jobHandler{jobService: jobService}
}



// FindJobs godoc
// @Summary Get all jobs
// @Description Get all jobs
// @Tags jobs
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /jobs [get]
func (h *jobHandler) FindJobs(ctx echo.Context) error {
	jobs, err := h.jobService.FindAllJob()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Succes get all jobs", jobs))
}


// FindSharedJobs godoc
// @Summary Get shared jobs
// @Description Get shared jobs
// @Tags jobs
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /jobs/shared [get]
func (h *jobHandler) FindSharedJobs(ctx echo.Context) error {
	dataUser, _ := ctx.Get("user").(*jwt.Token)
	claims := dataUser.Claims.(*token.JwtCustomClaims)
	jobs, err := h.jobService.FindSharedJobs(claims.ID)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK,response.SuccessResponse(http.StatusOK, "Succes Get Shared Jobs", jobs))
}

// FindAppliedJobs godoc
// @Summary Get applied jobs
// @Description Get applied jobs
// @Tags jobs
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /jobs/applied [get]
func (h *jobHandler) FindAppliedJobs(ctx echo.Context) error {
	dataUser, _ := ctx.Get("user").(*jwt.Token)
	claims := dataUser.Claims.(*token.JwtCustomClaims)
	jobs, err := h.jobService.FindAppliedJobs(claims.ID)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK,response.SuccessResponse(http.StatusOK, "Succes Get Applied Jobs", jobs))
}


// FindJobByID godoc
// @Summary Get job by ID
// @Description Get job by ID
// @Tags jobs
// @Produce json
// @Param id path string true "Job ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /jobs/{id} [get]
func (h *jobHandler) FindJobByID(ctx echo.Context) error {
	var input binder.JobFindByIDRequest

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	id := uuid.MustParse(input.ID)

	job, err := h.jobService.FindJobByID(id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Succes get job details", job))
}

// CreateJob godoc
// @Summary Create a new job
// @Description Create a new job
// @Tags jobs
// @Produce json
// @Security BearerAuth
// @Param job body binder.CreateJobRequest true "Create Job Request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /jobs [post]
func (h * jobHandler) CreateJob(ctx echo.Context) error {

	dataUser, _ := ctx.Get("user").(*jwt.Token)
	claims := dataUser.Claims.(*token.JwtCustomClaims)


	var input binder.CreateJobRequest
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	newJob := entity.NewJob(input.Title, input.Description, input.Company, input.Logo, input.Status, input.Salary, input.Location, input.CategoryID, claims.ID)

	job, err := h.jobService.CreateJob(newJob)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success create job", job))
}

// UpdateJob godoc
// @Summary Update a job
// @Description Update a job
// @Tags jobs
// @Produce json
// @Param id path string true "Job ID"
// @Param job body binder.UpdateJobRequest true "Update Job Request"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /jobs/{id} [put]
func (h *jobHandler) UpdateJob(ctx echo.Context) error {
   var input binder.UpdateJobRequest

   if err := ctx.Bind(&input); err != nil {
	   return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
   }

   updateJob := entity.UpdateJob(input.ID, input.Title, input.Description, input.Company, input.Logo, input.Status, input.Salary, input.Location)

   updatedJob, err := h.jobService.UpdateJob(updateJob)

   if err != nil {
	   return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
   }

   return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success update job", updatedJob))
}


// DeleteJob godoc
// @Summary Delete a job
// @Description Delete a job
// @Tags jobs
// @Produce json
// @Security BearerAuth
// @Param id path string true "Job ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /jobs/{id} [delete]
func (h *jobHandler) DeleteJob(ctx echo.Context) error {
	var input binder.DeleteJobRequest

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	id := uuid.MustParse(input.ID)

	isDeleted, err := h.jobService.DeleteJob(id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success delete job", isDeleted))
}
