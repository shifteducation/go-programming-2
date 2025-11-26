package main

import (
	"github.com/shifteducation/user-service/internal/repositories"
	"github.com/shifteducation/user-service/internal/router"
	"github.com/shifteducation/user-service/internal/services"
)

func main() {
	userPostgresRepository := repositories.UserPostgresRepository{}
	userService := services.NewUserService(userPostgresRepository)
	r := router.NewRouter(userService)
	r.Run()
}
