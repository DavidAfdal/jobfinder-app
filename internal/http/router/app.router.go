package router

import (
	"net/http"

	"github.com/DavidAfdal/workfinder/internal/http/handler"
	"github.com/DavidAfdal/workfinder/pkg/route"
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
		},
		{
			Methode: http.MethodGet,
			Path:    "/users",
			Handler: userHandler.FindAllUser,
		},
		{
			Methode: http.MethodPatch,
			Path: "/users/:id",
			Handler: userHandler.UpdateUser,
		},
		{
			Methode: http.MethodGet,
			Path:    "/jobs/shared",
			Handler: jobHandler.FindSharedJobs,
		},
		{
			Methode: http.MethodGet,
			Path: "/jobs/applied",
			Handler: jobHandler.FindAppliedJobs,
		},
		{
			Methode: http.MethodPost,
			Path: "/jobs",
			Handler: jobHandler.CreateJob,
		},
		{
			Methode: http.MethodPatch,
			Path: "/jobs/:id",
			Handler: jobHandler.UpdateJob,
		},
		{
			Methode: http.MethodPost,
			Path: "/jobs/:jobID/apply",
			Handler: jobApplicationHandler.ApplyJob,
		},
		{
			Methode: http.MethodGet,
			Path: "/jobs/:JobApplicantID/applications",
			Handler: jobApplicationHandler.FindJobApplicantsByID,
		},
		{
			Methode: http.MethodGet,
			Path: "/jobs/:JobApplicantID/withdraw",
			Handler: jobApplicationHandler.WithdrawnJobApplicants,
		},
		{
			Methode: http.MethodGet,
			Path: "/jobs/:JobApplicantID/approve",
			Handler: jobApplicationHandler.ApproveApplicant,
		},
		{
			Methode: http.MethodPost,
			Path: "/categories",
			Handler: categoryHandeler.CreateCategory,
		},
		{
			Methode: http.MethodPatch,
			Path: "/categories/:id",
			Handler: categoryHandeler.UpdateCategory,
		},
		{
			Methode: http.MethodDelete,
			Path: "/categories/:id",
			Handler: categoryHandeler.DeleteCategory,
		},
	}
}
