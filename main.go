package main

import (
	"database/sql"
	"fmt"

	"github.com/br4tech/go-to-mysql/adapter"
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

	adapter := adapter.NewMySQLAdapter(db)

	data, err := adapter.FindByID(1)
	if err != nil {
		panic(err)
	}

   fmt.Println(data)
}
