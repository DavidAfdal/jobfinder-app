package main

import (
	"github.com/DavidAfdal/workfinder/config"
	_ "github.com/DavidAfdal/workfinder/docs"
	"github.com/DavidAfdal/workfinder/internal/builder"
	"github.com/DavidAfdal/workfinder/pkg/cache"
	"github.com/DavidAfdal/workfinder/pkg/postgres"
	"github.com/DavidAfdal/workfinder/pkg/server"
	"github.com/DavidAfdal/workfinder/pkg/token"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg, err := config.NewConfig(".env")
	checkError(err)
	db, err := postgres.InitPostgres(&cfg.Postgres)
	checkError(err)

	tokenUseCase := token.NewTokenUseCase(cfg.JWT.SecretKey)
	redisDB := cache.InitCache(&cfg.Redis)

	publicRoutes := builder.BuildAppRoutes(db, tokenUseCase, redisDB)
	privateRoutes := builder.BuildPrivateAppRoutes(db, redisDB)


	srv:= server.NewServer("api", publicRoutes, privateRoutes, cfg.JWT.SecretKey)

	srv.Run()

}


func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
