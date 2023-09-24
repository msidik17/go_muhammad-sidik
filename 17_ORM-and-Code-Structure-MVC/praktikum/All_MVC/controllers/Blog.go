package controllers

import (
	"17_ORM-and-Code-Structure-MVC/praktikum/All_MVC/configs"
	"17_ORM-and-Code-Structure-MVC/praktikum/All_MVC/models"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// get all blogs
func GetBlogsController(c echo.Context) error {
	var blogs models.Blog
  
  
	if err := configs.DB.Find(&blogs).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success get all blogs",
	  "blogs":   blogs,
	})
  }

// get blog by id
func GetBlogController(c echo.Context) error {
	// Ambil parameter id dari URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
	  return echo.NewHTTPError(http.StatusBadRequest, "Invalid blog ID")
	}
  
	//Mengammbil blog Berdasarkan ID
	var blogs models.Blog
	if err := configs.DB.First(&blogs, id).Error; err != nil{
	  if gorm.IsRecordNotFoundError(err){
		  return echo.NewHTTPError(http.StatusNotFound, "blog Not Found")
	  }
	  return echo.NewHTTPError(http.StatusInternalServerError, "Database Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "Success get blog by ID",
	  "blog":    blogs,
	})
  }
  
  
  // create new blog
  func CreateBlogController(c echo.Context) error {
	blog := models.Blog{}
	c.Bind(&blog)
  
  
	if err := configs.DB.Save(&blog).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success create new blog",
	  "blog":    blog,
	})
  }
  
  
  // delete blog by id
  func DeleteBlogController(c echo.Context) error {
	// Ambil parameter id dari URL
	blogID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, "Invalid blog ID")
	}
  
	// Mencari blog Berdasarkan ID
	var blog models.Blog
	if err := configs.DB.First(&blog, blogID).Error; err != nil {
	  if gorm.IsRecordNotFoundError(err) {
		return echo.NewHTTPError(http.StatusNotFound, "blog not found")
	  }
	  return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}
  
	// Menghapus blog dari database
	if err := configs.DB.Delete(&blog).Error; err != nil {
	  return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete blog")
	}
  
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "Success delete blog",
	})
  }
  
  
  // update blog by id
  func UpdateBlogController(c echo.Context) error {
	  //Ambil parameter id dari URL
	blogID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, "Invalid blog ID")
	}
  
	// Mencari blog Berdasarkan ID
	var blog models.Blog
	if err := configs.DB.First(&blog, blogID).Error; err != nil {
	  if gorm.IsRecordNotFoundError(err) {
		return echo.NewHTTPError(http.StatusNotFound, "blog not found")
	  }
	  return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}
  
	// Bind Data blog Yang Baru
	updatedblog := new(models.Blog)
	if err := c.Bind(updatedblog); err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, "Invalid blog data")
	}
  
	// Update Data blog 
	blog.UserID = updatedblog.UserID
	blog.Title = updatedblog.Title
	blog.Content = updatedblog.Content
  
	if err := configs.DB.Save(&blog).Error; err != nil {
	  return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update blog")
	}
  
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "Success update blog",
	  "blog":    blog,
	})
  }
  