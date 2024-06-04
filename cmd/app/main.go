package main

import (
	"github.com/DavidAfdal/workfinder/config"
	"github.com/DavidAfdal/workfinder/internal/builder"
	"github.com/DavidAfdal/workfinder/pkg/cache"
	"github.com/DavidAfdal/workfinder/pkg/postgres"
	"github.com/DavidAfdal/workfinder/pkg/server"
	"github.com/DavidAfdal/workfinder/pkg/token"
)

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
