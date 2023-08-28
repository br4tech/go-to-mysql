package adapter

import (
	"database/sql"

	"github.com/br4tech/go-to-mysql/repository"
)

type MySQLAdapter struct {
	db *sql.DB
}

func NewMySQLAdapter(db *sql.DB) *MySQLAdapter {
	return &MySQLAdapter{db: db}
}

func (r *MySQLAdapter) FindByID(id int)(*repository.User, error){
	  // Consulta SQL para buscar um usuário pelo ID
    query := "SELECT id, nome, age FROM user_data WHERE id = ?"
    
    // Execute a consulta SQL
    row := r.db.QueryRow(query, id)
    
    // Crie uma estrutura para armazenar o resultado
    var user repository.User
    
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