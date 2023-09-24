package controllers

import (
	"17_ORM-and-Code-Structure-MVC/praktikum/All_MVC/configs"
	"17_ORM-and-Code-Structure-MVC/praktikum/All_MVC/models"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// get all users
func GetUsersController(c echo.Context) error {
	var users models.User
  
  
	if err := configs.DB.Find(&users).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success get all users",
	  "users":   users,
	})
  }

// get user by id
func GetUserController(c echo.Context) error {
	// Ambil parameter id dari URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
	  return echo.NewHTTPError(http.StatusBadRequest, "Invalid User ID")
	}
  
	//Mengammbil User Berdasarkan ID
	var users models.User
	if err := configs.DB.First(&users, id).Error; err != nil{
	  if gorm.IsRecordNotFoundError(err){
		  return echo.NewHTTPError(http.StatusNotFound, "User Not Found")
	  }
	  return echo.NewHTTPError(http.StatusInternalServerError, "Database Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "Success get user by ID",
	  "user":    users,
	})
  }
  
  
  // create new user
  func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
  
  
	if err := configs.DB.Save(&user).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success create new user",
	  "user":    user,
	})
  }
  
  
  // delete user by id
  func DeleteUserController(c echo.Context) error {
	// Ambil parameter id dari URL
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}
  
	// Mencari User Berdasarkan ID
	var user models.User
	if err := configs.DB.First(&user, userID).Error; err != nil {
	  if gorm.IsRecordNotFoundError(err) {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	  }
	  return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}
  
	// Menghapus User dari database
	if err := configs.DB.Delete(&user).Error; err != nil {
	  return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete user")
	}
  
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "Success delete user",
	})
  }
  
  
  // update user by id
  func UpdateUserController(c echo.Context) error {
	  //Ambil parameter id dari URL
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}
  
	// Mencari User Berdasarkan ID
	var users models.User
	if err := configs.DB.First(&users, userID).Error; err != nil {
	  if gorm.IsRecordNotFoundError(err) {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	  }
	  return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}
  
	// Bind Data User Yang Baru
	updatedUser := new(models.User)
	if err := c.Bind(updatedUser); err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, "Invalid user data")
	}
  
	// Update Data user 
	users.Name = updatedUser.Name
	users.Email = updatedUser.Email
	users.Password = updatedUser.Password
  
	if err := configs.DB.Save(&users).Error; err != nil {
	  return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update user")
	}
  
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "Success update user",
	  "user":    users,
	})
  }
  