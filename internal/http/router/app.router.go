package router

import (
	"net/http"

	"github.com/DavidAfdal/workfinder/internal/http/handler"
	"github.com/DavidAfdal/workfinder/pkg/route"
)

const (
	Applicant 	  = "Applicant"
	Client        = "Client"
)

var (
	allRoles   = []string{Applicant, Client}
	onlyApplicant  = []string{Applicant}
	onlyClient  = []string{Client}
)

func AppPublicRoutes(userHandler handler.UserHandler, jobHandler handler.JobHandler, categoryHandeler handler.CategoryHandler) []*route.Route {
	return []*route.Route{
		{
			Methode: http.MethodPost,
			Path:    "/login",
			Handler: userHandler.Login,
		},
		{
			Methode: http.MethodPost,
			Path:    "/register",
			Handler: userHandler.CreateUser,
		},
		{
			Methode: http.MethodGet,
			Path: "/jobs",
			Handler: jobHandler.FindJobs,
		},
		{
			Methode: http.MethodGet,
			Path: "/jobs/:id",
			Handler: jobHandler.FindJobByID,
		},
		{
			Methode: http.MethodGet,
			Path: "/users/:id",
			Handler: userHandler.FindByUserID,
		},
		{
			Methode: http.MethodGet,
			Path: "/categories",
			Handler: categoryHandeler.FindAllCategory,
		},
		{
			Methode: http.MethodGet,
			Path: "/categories/:id",
			Handler: categoryHandeler.FindCategoryByID,
		},
	}
}


func AppPrivateRoute(userHandler handler.UserHandler,  jobHandler handler.JobHandler, jobApplicationHandler handler.JobApplicantsHandler, categoryHandeler handler.CategoryHandler) []*route.Route {
	return []*route.Route{
		{
			Methode: http.MethodGet,
			Path:    "/profile",
			Handler: userHandler.ProfileUser,
			Role: allRoles,
		},
		{
			Methode: http.MethodGet,
			Path:    "/users",
			Handler: userHandler.FindAllUser,
			Role: allRoles,
		},
		{
			Methode: http.MethodPatch,
			Path: "/users/:id",
			Handler: userHandler.UpdateUser,
			Role: allRoles,
		},
		{
			Methode: http.MethodDelete,
			Path: "/users",
			Handler: userHandler.DeleteUser,
			Role: allRoles,
		},
		{
			Methode: http.MethodGet,
			Path:    "/jobs/shared",
			Handler: jobHandler.FindSharedJobs,
			Role: onlyClient,
		},
		{
			Methode: http.MethodGet,
			Path: "/jobs/applied",
			Handler: jobHandler.FindAppliedJobs,
			Role: onlyApplicant,
		},
		{
			Methode: http.MethodPost,
			Path: "/jobs",
			Handler: jobHandler.CreateJob,
			Role: onlyClient,
		},
		{
			Methode: http.MethodPatch,
			Path: "/jobs/:id",
			Handler: jobHandler.UpdateJob,
			Role: onlyClient,
		},
		{
			Methode: http.MethodDelete,
			Path: "/jobs/:id",
			Handler: jobHandler.DeleteJob,
			Role: onlyClient,
		},
		{
			Methode: http.MethodPost,
			Path: "/jobs/:jobID/apply",
			Handler: jobApplicationHandler.ApplyJob,
			Role: onlyApplicant,
		},
		{
			Methode: http.MethodGet,
			Path: "/jobs/:JobApplicantID/applications",
			Handler: jobApplicationHandler.FindJobApplicantsByID,
			Role: allRoles,
		},
		{
			Methode: http.MethodGet,
			Path: "/jobs/:JobApplicantID/withdraw",
			Handler: jobApplicationHandler.WithdrawnJobApplicants,
			Role: onlyApplicant,
		},
		{
			Methode: http.MethodGet,
			Path: "/jobs/:JobApplicantID/approve",
			Handler: jobApplicationHandler.ApproveApplicant,
			Role: onlyClient,
		},
		{
			Methode: http.MethodPost,
			Path: "/categories",
			Handler: categoryHandeler.CreateCategory,
			Role: onlyClient,
		},
		{
			Methode: http.MethodPatch,
			Path: "/categories/:id",
			Handler: categoryHandeler.UpdateCategory,
			Role: onlyClient,
		},
		{
			Methode: http.MethodDelete,
			Path: "/categories/:id",
			Handler: categoryHandeler.DeleteCategory,
			Role: onlyClient,
		},
	}
}
