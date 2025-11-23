package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/shifteducation/user-service/internal/dto"
	"github.com/shifteducation/user-service/internal/mocks"
	"github.com/shifteducation/user-service/internal/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetAllUsersOK(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	users := []models.User{
		{
			Id:        uuid.New(),
			FirstName: "John",
			LastName:  "Doe",
			Age:       25,
			Address:   nil,
		},
	}

	userService := mocks.NewMockUserService(mockCtrl)
	userService.EXPECT().
		GetAll(gomock.Any()).
		Return(users, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/users", nil)

	router := NewRouter(userService)
	router.engine.ServeHTTP(w, req)

	userJson, err := json.Marshal(users)
	if err != nil {
		t.Errorf("Error while marshalling json, %s", err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(userJson), w.Body.String())
}

func TestGetAllUsersServerError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userService := mocks.NewMockUserService(mockCtrl)
	userService.EXPECT().
		GetAll(gomock.Any()).
		Return(nil, errors.New("test error"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/users", nil)

	router := NewRouter(userService)
	router.engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGetUserByIdOK(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	id := uuid.New()
	user := &models.User{
		Id:        id,
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
		Address:   nil,
	}

	userService := mocks.NewMockUserService(mockCtrl)
	userService.EXPECT().
		GetById(gomock.Any(), id).
		Return(user, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/v1/users/%s", id.String()), nil)

	router := NewRouter(userService)
	router.engine.ServeHTTP(w, req)

	userJson, err := json.Marshal(user)
	if err != nil {
		t.Errorf("Error while marshalling json, %s", err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(userJson), w.Body.String())
}

func TestGetUserByIdBadRequest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userService := mocks.NewMockUserService(mockCtrl)

	wrongId := "wrongId"

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/v1/users/%s", wrongId), nil)

	router := NewRouter(userService)
	router.engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetUserByIdServerError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	id := uuid.New()

	userService := mocks.NewMockUserService(mockCtrl)
	userService.EXPECT().
		GetById(gomock.Any(), id).
		Return(nil, errors.New("test error"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/v1/users/%s", id.String()), nil)

	router := NewRouter(userService)
	router.engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestCreateUserOK(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	user := &models.User{
		Id:        uuid.New(),
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
		Address:   nil,
	}
	createUserRequest := dto.CreateUserRequest{
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
		Address:   nil,
	}
	createUserRequestJson, err := json.Marshal(createUserRequest)
	if err != nil {
		t.Errorf("Error while marshalling json, %s", err)
	}
	userService := mocks.NewMockUserService(mockCtrl)
	userService.EXPECT().
		Create(gomock.Any(), createUserRequest).
		Return(user, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/users", strings.NewReader(string(createUserRequestJson)))

	router := NewRouter(userService)
	router.engine.ServeHTTP(w, req)

	userJson, err := json.Marshal(user)
	if err != nil {
		t.Errorf("Error while marshalling json, %s", err)
	}

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, string(userJson), w.Body.String())
}

func TestCreateUserServerError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	createUserRequest := dto.CreateUserRequest{
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
		Address:   nil,
	}
	createUserRequestJson, err := json.Marshal(createUserRequest)
	if err != nil {
		t.Errorf("Error while marshalling json, %s", err)
	}
	userService := mocks.NewMockUserService(mockCtrl)
	userService.EXPECT().
		Create(gomock.Any(), createUserRequest).
		Return(nil, errors.New("test error"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/users", strings.NewReader(string(createUserRequestJson)))

	router := NewRouter(userService)
	router.engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestCreateUserBadRequest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	createUserRequest := dto.CreateUserRequest{}
	createUserRequestJson, err := json.Marshal(createUserRequest)
	if err != nil {
		t.Errorf("Error while marshalling json, %s", err)
	}
	userService := mocks.NewMockUserService(mockCtrl)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/users", strings.NewReader(string(createUserRequestJson)))

	router := NewRouter(userService)
	router.engine.ServeHTTP(w, req)

	errorMessage := `Key: 'CreateUserRequest.FirstName' Error:Field validation for 'FirstName' failed on the 'required' tag
Key: 'CreateUserRequest.LastName' Error:Field validation for 'LastName' failed on the 'required' tag
Key: 'CreateUserRequest.Age' Error:Field validation for 'Age' failed on the 'required' tag`

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, errorMessage, w.Body.String())
}
