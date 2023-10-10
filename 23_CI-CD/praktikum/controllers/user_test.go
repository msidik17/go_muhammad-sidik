package controllers_test

import (
	"docker/models"
	"docker/controllers"
	"docker/repositories"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	e := echo.New()
	mockUserRepo := &repositories.MockUserRepository{}
	uController := controllers.NewUserController(mockUserRepo)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	users := []models.User{}
	mockUserRepo.On("Find").Return(users, nil)

	err := uController.GetAllUsers(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	mockUserRepo.AssertExpectations(t)
}

func TestCreateUser(t *testing.T) {
	reqbody := `{
		"email": "muhamamdsidik@gmail.com",
		"password":"6789"}`

	mockUserRepo := &repositories.MockUserRepository{}
	uController := controllers.NewUserController(mockUserRepo)

	expectedUser := models.User{
		Email:    "muhamamdsidik@gmail.com",
		Password: "6789",
	}

	mockUserRepo.On("Create", expectedUser).Return(nil)

	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(reqbody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)

	err := uController.CreateUser(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	mockUserRepo.AssertExpectations(t)
}

func TestCreateUserfailed(t *testing.T) {
	reqbody := `{ "email": "muhammadsidik@gmail.com"}`

	mockUserRepo := &repositories.MockUserRepository{}
	uController := controllers.NewUserController(mockUserRepo)

	expectedUser := models.User{
		Email: "muhammadsidik@gmail.com",
	}

	mockUserRepo.On("Create", expectedUser).Return(errors.New("some other error"))

	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(reqbody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)

	if assert.NoError(t, uController.CreateUser(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	}

	mockUserRepo.AssertExpectations(t)
}
