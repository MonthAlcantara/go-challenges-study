package model

// User representa a entidade de usuário.
// Structs em Go são simples, sem getters/setters, e as tags facilitam (de)serialização.
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
