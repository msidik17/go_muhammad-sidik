package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
  Id    int    `json:"id" form:"id"`
  Name  string `json:"name" form:"name"`
  Email string `json:"email" form:"email"`
  Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
  return c.JSON(http.StatusOK, map[string]interface{}{
    "messages": "success get all users",
    "users":    users,
  })
}

// get user by id
func GetUserController(c echo.Context) error {
  // Ambil parameter id dari URL
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil{
	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"message" : "Invalid user ID",
	})
  }

  //Mencari User dengan ID
  var foundUser *User
  for _, user := range users {
	if user.Id == id {
		foundUser = &user
		break
	}
  }

  //Jika user tidak ditemukan, maka kembalikan dengan 404 not found
  if foundUser == nil{
    return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message" : "User Not Found",
	})
  }

  //Mengembalikan user yang ditemukan
  return c.JSON(http.StatusOK, map[string]interface{}{
	"message" : "Success get user",
	"user" : foundUser,
  })

}



// delete user by id
func DeleteUserController(c echo.Context) error {
   // Ambil parameter id dari URL
   id, err := strconv.Atoi(c.Param("id"))
   if err != nil{
	 return c.JSON(http.StatusBadRequest, map[string]interface{}{
		 "message" : "Invalid user ID",
	 })
   }

   //Menemukan indeks user menggunakan id 
   var foundIndex = -1
   for i, user := range users{
	if user.Id == id{
		foundIndex = i
		break
	}
   }

   //Jika user tidak ditemukan, maka kembalikan dengan 404 not found
  if foundIndex == -1{
    return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message" : "User Not Found",
	})
  }

  //Menghapus user dari slice users
  users = append(users[:foundIndex], users[foundIndex + 1:]...)

  return c.JSON(http.StatusOK, map[string]interface{}{
	"message" : "Success delete user",
  })
}



// update user by id
func UpdateUserController(c echo.Context) error {
    // Ambil parameter id dari URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
	  return c.JSON(http.StatusBadRequest, map[string]interface{}{
		  "message" : "Invalid user ID",
	  })
	}

	//Menemukan indeks user menggunakan id 
	var foundIndex = -1
	for i, user := range users{
	 if user.Id == id{
		 foundIndex = i
		 break
	 }
	}

	//Jika user tidak ditemukan, maka kembalikan dengan 404 not found
	if foundIndex == -1{
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message" : "User Not Found",
		})
	  }

	//Binding data user baru 
	updatedUser := User{}
	if err := c.Bind(&updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message" : "Invalid user ID",
		})
	}

	//Update user 
	users[foundIndex] = updatedUser

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "Success update user",
		"user" : updatedUser,
	  }) 
}

// create new user
func CreateUserController(c echo.Context) error {
  // binding data
  user := User{}
  c.Bind(&user)

  if len(users) == 0 {
    user.Id = 1
  } else {
    newId := users[len(users)-1].Id + 1
    user.Id = newId
  }
  users = append(users, user)
  return c.JSON(http.StatusOK, map[string]interface{}{
    "messages": "success create user",
    "user":     user,
  })
}

// ---------------------------------------------------
func main() {
  e := echo.New()
  // routing with query parameter
  e.GET("/users", GetUsersController)
  e.POST("/users", CreateUserController)
  e.GET("/users/:id", GetUserController)
  e.DELETE("/users/:id", DeleteUserController)
  e.PUT("/users/:id", UpdateUserController)

  // start the server, and log if it fails
  e.Logger.Fatal(e.Start(":8000"))
}