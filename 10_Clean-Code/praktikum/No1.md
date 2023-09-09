package main

type user struct {
id int
username int
password int
}

type userservice struct {
t []user
}

func (u userservice) getallusers() []user {
return u.t
}

func (u userservice) getuserbyid(id int) user {
for \_, r := range u.t {
if id == r.id {
return r
}
}

    return user{}

}

pada code diatas banyak penamaan dan memilihan type data yang tidak seharus nya seperti :

- type data username dan password pada struktur data user lebih baik menggunakan 'string'
- menggunakan singkatan yang jelas seperti 't' pada struktur data userservis lebih baik menggunakan 'users'
- singkatan 'u' pada function userservis membuat sangat sulit untuk memahami apa arti dari singakatan tersebut, lebih baik menggunakan 'us' untuk singkatan dari 'userservice'

berikut adalah revisi dari code diatas yang telah menggunakan standar progamming.

package main

type User struct {
ID int
Name string
Password string
}

type UserService struct {
Users []User
}

func (us UserService) GetAllUsers() []User {
return us.Users
}

func (us UserService) GetUserByID(userID int) User {
for \_, u := range us.Users {
if userID == u.UserID {
return u
}
}

    return User{}

}
