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
	// DeleteJob(ctx echo.Context) error
}

type jobHandler struct {
	jobService service.JobService
}


func NewJobHandler(jobService service.JobService) JobHandler {
	return &jobHandler{jobService: jobService}
}


func (h *jobHandler) FindJobs(ctx echo.Context) error {
	jobs, err := h.jobService.FindAllJob()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Succes get all jobs", jobs))
}

func (h *jobHandler) FindSharedJobs(ctx echo.Context) error {
	dataUser, _ := ctx.Get("user").(*jwt.Token)
	claims := dataUser.Claims.(*token.JwtCustomClaims)
	jobs, err := h.jobService.FindSharedJobs(claims.ID)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK,response.SuccessResponse(http.StatusOK, "Succes Get Shared Jobs", jobs))
}
func (h *jobHandler) FindAppliedJobs(ctx echo.Context) error {
	dataUser, _ := ctx.Get("user").(*jwt.Token)
	claims := dataUser.Claims.(*token.JwtCustomClaims)
	jobs, err := h.jobService.FindAppliedJobs(claims.ID)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK,response.SuccessResponse(http.StatusOK, "Succes Get Applied Jobs", jobs))
}

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
