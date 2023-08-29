package repository

import (
	"database/sql"
	"testing"

	"github.com/br4tech/go-to-mysql/adapter"
)

func TestUserRepository_FindByID(t *testing.T) {
    // Crie um banco de dados em memória para o teste
    db, err := sql.Open("sqlite3", ":memory:")
    if err != nil {
        t.Fatalf("Erro ao criar o banco de dados em memória: %v", err)
    }
    defer db.Close()

    // Crie uma instância do MockMySQLAdapter usando o banco de dados em memória
    mockAdapter := adapter.NewMySQLAdapter(db)

    // Crie uma instância do UserRepository usando o mockAdapter
    repo := NewUserRepository(mockAdapter)

    // Chame a função FindByID
    user, err := repo.FindByID(1)

    // Verifique se não há erros
    if err != nil {
        t.Fatalf("Erro inesperado: %v", err)
    }

    // Verifique se o usuário retornado está correto
    expectedUser := &User{
        ID:   1,
        Name: "Nome do Usuário",
        Age:  30,
    }

    if user.ID != expectedUser.ID || user.Name != expectedUser.Name || user.Age != expectedUser.Age {
        t.Errorf("Resultado incorreto. Esperado %v, mas obteve %v", expectedUser, user)
    }
}

// MockMySQLAdapter é uma implementação falsa da interface adapter.MySQLAdapter
type MockMySQLAdapter struct {
    db *sql.DB
}

// GetDB retorna o banco de dados simulado
func (m *MockMySQLAdapter) GetDB() *sql.DB {
    return m.db
}

