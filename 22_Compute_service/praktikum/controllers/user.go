package controllers

import (
	"docker/dto"
	"docker/middleware"
	"docker/models"
	"docker/repositories"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	GetAllUsers(c echo.Context) error
	CreateUser(c echo.Context) error
}

type userController struct{
	userRepo repositories.UserRepository
}

func NewUserController(uRepo repositories.UserRepository) *userController{
	return &userController{
		userRepo: uRepo,
	}
}

func (u *userController) GetAllUsers(c echo.Context) error{
	users, err := u.userRepo.Find()
	if err != nil{
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error" : err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"Data" : users,
	})
}

func (u *userController) CreateUser(c echo.Context) error{
	var user models.User
	err := c.Bind(&user)
	
	if err != nil{
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error" : err.Error(),
		})
	}

	err = u.userRepo.Create(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": fmt.Sprintf("Failed to create user: %v", err),
		})
	}

	token, err := middleware.CreateToken(user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": fmt.Sprintf("failed to create token: %v", err),
		})
	}
	userResult := dto.DTOUsers{
		Email: user.Email,
		Token: token,
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": userResult,
	})
}