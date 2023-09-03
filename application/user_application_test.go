package app

import (
	"errors"
	"testing"

	"github.com/br4tech/go-to-mysql/repository"
)

// Implementação de um UserService de mock para fins de teste.
type MockUserService struct{}

func (m *MockUserService) GetUserByID(id int) (*repository.User, error) {
	// Simule uma chamada ao serviço que retorna um usuário ou um erro, dependendo do valor de "id".
	if id == 1 {
		return &repository.User{ID: 1, Name: "Alice", Age: 30}, nil
	}
	return nil, errors.New("Usuário não encontrado")
}

func TestApplication_GetUserByID(t *testing.T) {
	// Crie uma instância do seu Application com o UserService de mock.
	application := NewApplication(&MockUserService{})

	// Teste o cenário em que o usuário é encontrado.
	user, err := application.GetUserByID(1)
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
	user, err = application.GetUserByID(2)
	if err == nil {
		t.Error("Esperava um erro, mas não ocorreu")
	}
	if user != nil {
		t.Error("O usuário deveria ser nulo")
	}
}
