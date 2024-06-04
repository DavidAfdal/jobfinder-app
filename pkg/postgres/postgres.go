package postgres

import (
	"fmt"

	"github.com/DavidAfdal/workfinder/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitPostgres(config *config.PostgresConfig) (*gorm.DB, error) {
   dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.Database)

   db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
	  Logger: logger.Default.LogMode(logger.Info),
   })
   if err != nil {
	return db, err
   }

   return db, nil
}
