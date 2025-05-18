package main

import (
	"go-bd-study/internal/db"
	"go-bd-study/internal/handler"
	"go-bd-study/internal/repository"
	"log"
	"net/http"
)

func main() {
	// Cria uma conexão com o banco de dados SQLite (arquivo teste.db)
	// Em Java, seria parecido com usar JDBC para abrir uma conexão.
	database := db.New("./teste.db")
	defer database.Close() // Garante que o banco será fechado ao final da função (similar ao try-with-resources do Java)

	// Cria o repositório de usuários, injetando a conexão do banco
	// Em Java, seria como instanciar um DAO passando a Connection.
	repo := &repository.UserRepository{DB: database}

	// Cria o handler de usuário, injetando o repositório
	// Em Java, seria como passar o DAO para um Controller.
	userHandler := &handler.UserHandler{Repo: repo}

	// Define a rota "/users" e um handler para ela
	// Em Java, seria como mapear um endpoint com @WebServlet ou @RequestMapping.
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method { // Verifica o método HTTP da requisição
		case http.MethodGet:
			userHandler.ListUsers(w, r) // GET: lista usuários
		case http.MethodPost:
			userHandler.CreateUser(w, r) // POST: cria usuário
		default:
			// Retorna erro 405 se o método não for suportado
			http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		}
	})

	// Loga que o servidor foi iniciado
	// Em Java, seria como System.out.println ou Logger.info
	log.Println("Servidor iniciado em :8080")

	// Inicia o servidor HTTP na porta 8080
	// Se der erro, log.Fatal imprime o erro e encerra o programa
	// Em Java, seria como iniciar um servidor com Jetty/Tomcat e tratar exceções.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
