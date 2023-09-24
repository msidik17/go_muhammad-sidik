package controllers

import (
	"17_ORM-and-Code-Structure-MVC/praktikum/All_MVC/configs"
	"17_ORM-and-Code-Structure-MVC/praktikum/All_MVC/models"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// get all books
func GetBooksController(c echo.Context) error {
	var books models.Book
  
  
	if err := configs.DB.Find(&books).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success get all books",
	  "books":   books,
	})
  }

// get book by id
func GetBookController(c echo.Context) error {
	// Ambil parameter id dari URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
	  return echo.NewHTTPError(http.StatusBadRequest, "Invalid book ID")
	}
  
	//Mengammbil book Berdasarkan ID
	var books models.Book
	if err := configs.DB.First(&books, id).Error; err != nil{
	  if gorm.IsRecordNotFoundError(err){
		  return echo.NewHTTPError(http.StatusNotFound, "book Not Found")
	  }
	  return echo.NewHTTPError(http.StatusInternalServerError, "Database Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "Success get book by ID",
	  "book":    books,
	})
  }
  
  
  // create new book
  func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)
  
  
	if err := configs.DB.Save(&book).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success create new book",
	  "book":    book,
	})
  }
  
  
  // delete book by id
  func DeleteBookController(c echo.Context) error {
	// Ambil parameter id dari URL
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, "Invalid book ID")
	}
  
	// Mencari book Berdasarkan ID
	var book models.Book
	if err := configs.DB.First(&book, bookID).Error; err != nil {
	  if gorm.IsRecordNotFoundError(err) {
		return echo.NewHTTPError(http.StatusNotFound, "book not found")
	  }
	  return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}
  
	// Menghapus book dari database
	if err := configs.DB.Delete(&book).Error; err != nil {
	  return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete book")
	}
  
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "Success delete book",
	})
  }
  
  
  // update book by id
  func UpdateBookController(c echo.Context) error {
	  //Ambil parameter id dari URL
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, "Invalid book ID")
	}
  
	// Mencari book Berdasarkan ID
	var book models.Book
	if err := configs.DB.First(&book, bookID).Error; err != nil {
	  if gorm.IsRecordNotFoundError(err) {
		return echo.NewHTTPError(http.StatusNotFound, "book not found")
	  }
	  return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}
  
	// Bind Data book Yang Baru
	updatedBook := new(models.Book)
	if err := c.Bind(updatedBook); err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, "Invalid book data")
	}
  
	// Update Data book 
	book.Title = updatedBook.Title
	book.Author = updatedBook.Author
	book.Publisher = updatedBook.Publisher
  
	if err := configs.DB.Save(&book).Error; err != nil {
	  return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update book")
	}
  
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "Success update book",
	  "book":    book,
	})
  }
  