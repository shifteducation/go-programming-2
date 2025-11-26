package main

import (
	"github.com/shifteducation/user-service/internal/configs"
	"github.com/shifteducation/user-service/internal/repositories"
	"github.com/shifteducation/user-service/internal/router"
	"github.com/shifteducation/user-service/internal/services"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	var config configs.AppConfig
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode config into struct, %v", err)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: config.DB.DSN,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(config.DB.LogLevel)),
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

	r := router.NewRouter(userService, config.App.Port)
	r.Run()
}
