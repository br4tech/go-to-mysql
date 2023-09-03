package service

import (
	"errors"
	"testing"

	"github.com/br4tech/go-to-mysql/repository"
)

// Implementação de um UserRepository de mock para fins de teste.
type MockUserRepository struct{}

func (m *MockUserRepository) FindByID(id int) (*repository.User, error) {
	// Simule uma consulta que retorna um usuário ou um erro, dependendo do valor de "id".
	if id == 1 {
		return &repository.User{ID: 1, Name: "Alice", Age: 30}, nil
	}
	return nil, errors.New("Usuário não encontrado")
}

func TestUserService_GetUserByID(t *testing.T) {
	// Crie uma instância do seu serviço com o UserRepository de mock.
	userService := NewUserService(&MockUserRepository{})

	// Teste o cenário em que o usuário é encontrado.
	user, err := userService.GetUserByID(1)
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}
	if user == nil {
		t.Error("O usuário não deveria ser nulo")
	}
	if user.Name != "Alice" {
		t.Errorf("Nome do usuário incorreto. Esperado 'Alice', mas obteve '%s'", user.Name)
	}

	// Teste o cenário em que o usuário não é encontrado.
	user, err = userService.GetUserByID(2)
	if err == nil {
		t.Error("Esperava um erro, mas não ocorreu")
	}
	if user != nil {
		t.Error("O usuário deveria ser nulo")
	}
}
