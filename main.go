package main

import (
	"database/sql"
	"fmt"

	"github.com/br4tech/go-to-mysql/adapter"
	"github.com/br4tech/go-to-mysql/repository"
	"github.com/br4tech/go-to-mysql/service"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
    // Configurar a conexão com o banco de dados MySQL
    db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/users")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // Criar uma instância do adaptador MySQL
    mysqlAdapter := adapter.NewMySQLAdapter(db)

    // Criar uma instância do repositório MySQLUserRepository, passando o adaptador MySQL
    userRepository := repository.NewUserRepository(mysqlAdapter)

    // Criar uma instância do serviço de usuário, passando o repositório de usuário
    userService := service.NewUserService(userRepository)

    // Exemplo de uso do serviço para buscar um usuário pelo ID
    user, err := userService.GetUserByID(1)
    if err != nil {
        fmt.Println("Erro ao buscar o usuário:", err)
        return
    }

    fmt.Printf("ID: %d\nNome: %s\nIdade: %s\n", user.ID, user.Name, user.Age)
}