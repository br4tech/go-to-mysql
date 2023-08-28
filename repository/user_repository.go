package repository

type UserRepository interface {
    FindByID(id int) (*User, error)
}

type User struct {
    ID       int
    Name string
    Age  int
}