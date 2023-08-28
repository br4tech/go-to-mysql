package app

import (
	"github.com/br4tech/go-to-mysql/repository"
	"github.com/br4tech/go-to-mysql/service"
)

type Application struct {
    userService service.IUserService
}

func NewApplication(userService service.IUserService) *Application {
    return &Application{
        userService: userService,
    }
}

func (app *Application) GetUserByID(id int) (*repository.User, error) {
    return app.userService.GetUserByID(id)
}

// Outros métodos de aplicação conforme necessário...
