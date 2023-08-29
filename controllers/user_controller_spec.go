package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/br4tech/go-to-mysql/repository"
	assert "github.com/stretchr/testify/assert"
)

// Implementação de um MockApplication para fins de teste.
type MockApplication struct{}

func (m *MockApplication) GetUserByID(id int) (*repository.User, error) {
	// Simule uma chamada ao serviço que retorna um usuário ou um erro, dependendo do valor de "id".
	if id == 1 {
		return &repository.User{ID: 1, Name: "Alice", Age: 30}, nil
	}
	return nil, nil // Simula um usuário não encontrado
}

func TestUserController_GetUserByIDHandler(t *testing.T) {
	// Crie uma instância do UserController com o MockApplication.
	userController := NewUserController(&MockApplication{})

	// Crie uma solicitação HTTP simulada para buscar o usuário com ID 1.
	req, err := http.NewRequest("GET", "/user/1", nil)
	assert.NoError(t, err)

	// Crie um ResponseWriter simulado para capturar a resposta.
	w := httptest.NewRecorder()

	// Chame a função GetUserByIDHandler do controlador.
	userController.GetUserByIDHandler(w, req)

	// Verifique a resposta HTTP.
	assert.Equal(t, http.StatusOK, w.Code)

	// Verifique a resposta do corpo. Dependendo de como você responde em GetUserByIDHandler, ajuste isso.
	expectedResponse := "ID: 1\nNome Alice\nIdade: 30\n"
	assert.Equal(t, expectedResponse, w.Body.String())

	// Teste o cenário em que o usuário não é encontrado (ID 2).
	req, err = http.NewRequest("GET", "/user/2", nil)
	assert.NoError(t, err)

	w = httptest.NewRecorder()
	userController.GetUserByIDHandler(w, req)

	// Verifique a resposta HTTP.
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	// Verifique a resposta do corpo.
	expectedErrorResponse := "Erro ao buscar o usuário: Usuário não encontrado"
	assert.Equal(t, expectedErrorResponse, w.Body.String())
}
