package repository

import (
	"database/sql"

	"github.com/br4tech/go-to-mysql/adapter"
)

type IUserRepository interface {
	 FindByID(id int) (*User, error)
}

type UserRepository struct {
	adapter *adapter.MySQLAdapter
}

func NewUserRepository(adapter *adapter.MySQLAdapter) IUserRepository {
    return &UserRepository{adapter: adapter}
}

func (r *UserRepository) FindByID(id int)(*User, error){
	  // Consulta SQL para buscar um usuário pelo ID
    query := "SELECT id, nome, age FROM user_data WHERE id = ?"
    
    // Execute a consulta SQL
    row := r.adapter.GetDB().QueryRow(query, id)
    
    // Crie uma estrutura para armazenar o resultado
    var user User
    
    // Faça o scan dos resultados da consulta para a estrutura
    err := row.Scan(&user.ID, &user.Name, &user.Age)
    if err != nil {
        if err == sql.ErrNoRows {
            // Não encontrou nenhum registro com o ID fornecido
            panic(err)
        }
        // Houve um erro durante a consulta
        return nil, err
    }
    
    return &user, nil
}

type User struct {
    ID       int
    Name string
    Age  int
}