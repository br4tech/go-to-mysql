package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// db esta a nivel de classe
var(
	db *sql.DB
)

type User struct {
	Name string
	Age int
}

func main(){
	var err error

	db, err = sql.Open("mysql", "root:admin@tcp(localhost:3306)/users")
	if err != nil {
		panic(err)
	}

	fmt.Println("Conectado")
	
	user := User {
		Name: "Joao Lucas",
		Age: 1,
	}

	if insertError := inserUser(user); insertError != nil {
		panic(err)
	}

	users, err := getAllUser()
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		fmt.Println(*user)
	}
}

func inserUser(user User) error {
	_, err := db.Exec(fmt.Sprintf("INSERT INTO user_data VALUES('%s', %d)", user.Name, user.Age))
	if err != nil {
		return err
	}

	fmt.Println("Usuario cadastrado com sucesso!")
	return nil
}

func getAllUser()([]*User, error) {
	res, err := db.Query("SELECT * FROM  user_data")
	if err != nil {
		return nil, err
	}
  
	users := []*User{}
	for res.Next(){
		var user User

		if err := res.Scan(&user.Name, &user.Age); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}