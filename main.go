package main

import (
	"github.com/shifteducation/user-service/internal/router"
	"github.com/shifteducation/user-service/internal/services"
)

func main() {
	userService := services.NewUserService(nil)
	r := router.NewRouter(userService)
	r.Run()
}
