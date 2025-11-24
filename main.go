package main

import (
	"github.com/shifteducation/user-service/internal/repositories"
	"github.com/shifteducation/user-service/internal/router"
	"github.com/shifteducation/user-service/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func main() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "postgres://user:password@localhost:5432/user-service",
		//PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Couldn't get sqlDB, err: %s", err.Error())
	}
	sqlDB.SetMaxOpenConns(10)

	userPostgresRepository := repositories.NewUserPostgresRepository(db)
	userService := services.NewUserService(userPostgresRepository)

	r := router.NewRouter(userService)
	r.Run()
}
