package repository

import (
	"database/sql"
	"go-bd-study/internal/model"
)

// UserRepository encapsula operações de acesso a dados.
// Em Go, structs são usadas para agrupar dados e métodos relacionados.
// Em Java, seria parecido com uma classe DAO (Data Access Object) que tem um atributo Connection.
type UserRepository struct {
	DB *sql.DB // Ponteiro para a conexão com o banco de dados (em Java seria java.sql.Connection)
}

// FindAll busca todos os usuários no banco de dados.
// Em Java, seria um método como List<User> findAll() em um DAO.
func (r *UserRepository) FindAll() ([]model.User, error) {
	rows, err := r.DB.Query("SELECT id, name FROM users") // Executa a query SQL (em Java: PreparedStatement + executeQuery)
	if err != nil {
		return nil, err // Retorna erro se a query falhar (em Java, lançaria uma exceção)
	}
	defer rows.Close() // Garante que os recursos do banco serão liberados (similar ao try-with-resources do Java)

	var users []model.User // Slice para armazenar os usuários (em Java seria List<User>)
	for rows.Next() {      // Itera sobre os resultados (em Java: while(rs.next()))
		var u model.User
		if err := rows.Scan(&u.ID, &u.Name); err != nil { // Mapeia as colunas para o struct (em Java: rs.getInt, rs.getString)
			return nil, err
		}
		users = append(users, u) // Adiciona o usuário ao slice (em Java: list.add(u))
	}
	return users, rows.Err() // Retorna a lista e possíveis erros de iteração
}

// Create insere um novo usuário no banco de dados.
// Em Java, seria void create(String name) em um DAO.
func (r *UserRepository) Create(name string) error {
	_, err := r.DB.Exec("INSERT INTO users(name) VALUES(?)", name) // Executa o insert (em Java: PreparedStatement + executeUpdate)
	return err                                                     // Retorna erro se houver (em Java, lançaria exceção)
}
