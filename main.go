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

	user := User {
		Name: "Joao Lucas",
		Age: 1,
	}

	if insertError := inserUser(user); insertError != nil {
		panic(err)
	}

	fmt.Println("Conectado")
}

func inserUser(user User) error {
	_, err := db.Exec(fmt.Sprintf("INSERT INTO user_data VALUES('%s', %d)", user.Name, user.Age))
	if err != nil {
		return err
	}

	fmt.Println("Usuario cadastrado com sucesso!")
	return nil
}