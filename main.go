package main

import (
	"database/sql"
	"net/http"

	"github.com/br4tech/go-to-mysql/adapter"
	app "github.com/br4tech/go-to-mysql/application"
	"github.com/br4tech/go-to-mysql/controllers"
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

    // Criar uma instância do repositório IUserRepository, passando o adaptador MySQL
    userRepository := repository.NewUserRepository(mysqlAdapter)

    // Criar uma instância do serviço IUserService, passando o repositório IUserRepository
    userService := service.NewUserService(userRepository)

    // Criar uma instância da aplicação, passando o serviço IUserService
    application := app.NewApplication(userService)

    // Criar uma instância do controlador de usuário, passando a aplicação
    userController := controllers.NewUserController(application)

    // Configurar rotas HTTP para manipuladores
    http.HandleFunc("/user", userController.GetUserByIDHandler)

    // Iniciar o servidor HTTP
    http.ListenAndServe(":8080", nil)
}