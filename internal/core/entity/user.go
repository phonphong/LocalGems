package entity

type User struct {
    ID       int    `json:"id"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
    Name     string `json:"name" binding:"required"`
}

type Credentials struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}