package router

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shifteducation/user-service/internal/dto"
	"github.com/shifteducation/user-service/internal/interfaces"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const paramId = "id"

type Router struct {
	userService interfaces.UserService
}

func NewRouter(userService interfaces.UserService) Router {
	return Router{
		userService: userService,
	}
}

func (r Router) Run() {
	router := gin.Default()

	routerGroup := router.Group("/api/v1")
	routerGroup.POST("/users", r.createUser)
	routerGroup.GET("/users", r.getAllUsers)
	routerGroup.GET("/users/:id", r.getUserById)
	routerGroup.PATCH("/users/:id", r.updateUser)
	routerGroup.DELETE("/users/:id", r.deleteUser)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

func (r Router) createUser(c *gin.Context) {
	userDto := dto.CreateUserRequest{}
	err := c.BindJSON(&userDto)
	if err != nil {
		log.Printf("Error while parsing user dto: %s\n", err.Error())
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	user, err := r.userService.Create(c, userDto)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (r Router) getAllUsers(c *gin.Context) {
	users, err := r.userService.GetAll(c)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, users)
}

func (r Router) getUserById(c *gin.Context) {
	paramId := c.Param(paramId)
	id, err := uuid.Parse(paramId)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	user, err := r.userService.GetById(c, id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (r Router) updateUser(c *gin.Context) {

}

func (r Router) deleteUser(c *gin.Context) {

}
