package builder

import (
	"github.com/DavidAfdal/workfinder/internal/http/handler"
	"github.com/DavidAfdal/workfinder/internal/http/router"
	"github.com/DavidAfdal/workfinder/internal/repository"
	"github.com/DavidAfdal/workfinder/internal/service"
	"github.com/DavidAfdal/workfinder/pkg/cache"
	"github.com/DavidAfdal/workfinder/pkg/route"
	"github.com/DavidAfdal/workfinder/pkg/token"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func BuildAppRoutes(db *gorm.DB, token token.TokenUseCase, redis *redis.Client ) []*route.Route {
	cahceable := cache.NewCacheable(redis)
	userRepository := repository.NewUserRepository(db, cahceable)
	userService := service.NewUserService(userRepository, token)
	userHandler := handler.NewUserHandler(userService)

	jobRepository := repository.NewJobRepository(db, cahceable)
	jobService := service.NewJobService(jobRepository)
	jobHandler := handler.NewJobHandler(jobService)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	return router.AppPublicRoutes(userHandler, jobHandler, categoryHandler)
}

func BuildPrivateAppRoutes(db *gorm.DB, redis *redis.Client) []*route.Route {
	cahceable := cache.NewCacheable(redis)
	userRepository := repository.NewUserRepository(db, cahceable)
	userService := service.NewUserService(userRepository, nil)
	userHandler := handler.NewUserHandler(userService)


	jobRepository := repository.NewJobRepository(db, cahceable)
	jobService := service.NewJobService(jobRepository)
	jobHandler := handler.NewJobHandler(jobService)

	jobApplicantsRepo := repository.NewJobApplicantsRepository(db)
	jobApplicantsService := service.NewJobApplicantService(jobApplicantsRepo, jobRepository)
	jobApplicantHandler := handler.NewJobApplicantsHandler(jobApplicantsService)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	return router.AppPrivateRoute(userHandler, jobHandler, jobApplicantHandler, categoryHandler)
}
