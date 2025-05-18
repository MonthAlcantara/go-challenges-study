package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// New abre (ou cria) o banco de dados SQLite e retorna a conexão.
// O padrão é retornar *sql.DB, que é thread-safe e gerencia múltiplas conexões.
func New(dataSourceName string) *sql.DB {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	// Criação da tabela se não existir.
	sqlStmt := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL
    );`
	if _, err := db.Exec(sqlStmt); err != nil {
		log.Fatal(err)
	}
	return db
}
