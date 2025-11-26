package router

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shifteducation/user-service/internal/custom_errors"
	"github.com/shifteducation/user-service/internal/dto"
	"github.com/shifteducation/user-service/internal/interfaces"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const paramId = "id"

type Router struct {
	engine      *gin.Engine
	userService interfaces.UserService
	port        int
}

func NewRouter(userService interfaces.UserService, port int) Router {
	router := Router{
		engine:      gin.Default(),
		userService: userService,
		port:        port,
	}

	routerGroup := router.engine.Group("/api/v1")
	routerGroup.POST("/users", router.createUser)
	routerGroup.GET("/users", router.getAllUsers)
	routerGroup.GET("/users/:id", router.getUserById)
	routerGroup.PATCH("/users/:id", router.updateUser)
	routerGroup.DELETE("/users/:id", router.deleteUser)

	return router
}

func (r Router) Run() {
	srv := &http.Server{
		Addr:    strings.Join([]string{":", strconv.Itoa(r.port)}, ""),
		Handler: r.engine,
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
	var notFoundError custom_errors.UserNotFoundError
	if errors.As(err, &notFoundError) {
		c.String(http.StatusNotFound, notFoundError.Error())
		return
	}
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (r Router) updateUser(c *gin.Context) {
	// changing Address
	// 204
	// 500
}

func (r Router) deleteUser(c *gin.Context) {
	// 204
	// 500
}
