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

// @title Job Finder App API
// @version 1.0
// @description This is a Rest API job finder server.
// @termsOfService http://swagger.io/terms/

// @contact.name David Afdal Kaizar Mutahadi
// @contact.url https://davidafdal.github.io/web-portofolio/
// @contact.email davidafdal7@gmail.com

// @host 103.157.26.216:8080
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
