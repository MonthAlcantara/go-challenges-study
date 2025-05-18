package handler

import (
	"encoding/json"
	"go-bd-study/internal/repository"
	"net/http"
)

// UserHandler lida com requisições HTTP relacionadas a usuários.
// Em Go, structs são usadas para agrupar dependências e métodos relacionados.
// Em Java, seria parecido com um Controller que recebe um DAO via injeção de dependência.
type UserHandler struct {
	Repo *repository.UserRepository // Referência ao repositório de usuários (em Java: private UserRepository repo;)
}

// GET /users
// Handler para listar usuários. Em Java, seria um método anotado com @GetMapping("/users") em um Controller.
func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Repo.FindAll() // Busca todos os usuários no repositório (em Java: repo.findAll())
	if err != nil {
		http.Error(w, "Erro ao buscar usuários", http.StatusInternalServerError) // Retorna erro 500 se falhar (em Java: response.sendError(500))
		return
	}
	json.NewEncoder(w).Encode(users) // Serializa a lista de usuários em JSON e escreve na resposta (em Java: new ObjectMapper().writeValue(response.getWriter(), users))
}

// POST /users
// Handler para criar usuário. Em Java, seria um método anotado com @PostMapping("/users") em um Controller.
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"` // Estrutura para receber o JSON do corpo da requisição (em Java: classe DTO ou @RequestBody)
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { // Decodifica o JSON recebido (em Java: new ObjectMapper().readValue(request.getInputStream(), Req.class))
		http.Error(w, "JSON inválido", http.StatusBadRequest) // Retorna erro 400 se o JSON for inválido (em Java: response.sendError(400))
		return
	}
	if err := h.Repo.Create(req.Name); err != nil { // Cria o usuário no repositório (em Java: repo.create(req.getName()))
		http.Error(w, "Erro ao criar usuário", http.StatusInternalServerError) // Retorna erro 500 se falhar
		return
	}
	w.WriteHeader(http.StatusCreated) // Retorna status 201 Created (em Java: response.setStatus(201))
}
