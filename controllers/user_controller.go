package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	app "github.com/br4tech/go-to-mysql/application"
)

type UserController struct {
    app *app.Application
}

func NewUserController(app *app.Application) *UserController {
    return &UserController{app: app}
}

func (c *UserController) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
    // Extrair o ID do usuário dos parâmetros da solicitação (por exemplo, de uma URL)
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    // Chamar o método de aplicação para buscar o usuário pelo ID
    user, err := c.app.GetUserByID(id)
    if err != nil {
        http.Error(w, fmt.Sprintf("Erro ao buscar o usuário: %v", err), http.StatusInternalServerError)
        return
    }

    // Responder com os dados do usuário em JSON ou HTML, por exemplo
    // Aqui, estamos apenas imprimindo na saída padrão como exemplo
    fmt.Printf("ID: %d\nNome %s\nIdade: %s\n", user.ID, user.Name, user.Age)
}