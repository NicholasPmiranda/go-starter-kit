package clientRequest

import (
	"github.com/jackc/pgx/v5/pgtype"
)

// CreateClientRequest representa os dados necessários para criar um cliente
// com validações do gin-gonic
type CreateClientRequest struct {
	Name    string      `json:"name" binding:"required,min=3,max=100"`
	Email   string      `json:"email" binding:"required,email"`
	Phone   pgtype.Text `json:"phone" binding:"omitempty"`
	Address pgtype.Text `json:"address" binding:"omitempty"`
}

// UpdateClientRequest representa os dados necessários para atualizar um cliente
// com validações do gin-gonic
type UpdateClientRequest struct {
	Name    string      `json:"name" binding:"required,min=3,max=100"`
	Email   string      `json:"email" binding:"required,email"`
	Phone   pgtype.Text `json:"phone" binding:"omitempty"`
	Address pgtype.Text `json:"address" binding:"omitempty"`
}

// ToCreateClientParams converte a request para o formato esperado pelo sqlc
func (r *CreateClientRequest) ToCreateClientParams() interface{} {
	return struct {
		Name    string      `json:"name"`
		Email   string      `json:"email"`
		Phone   pgtype.Text `json:"phone"`
		Address pgtype.Text `json:"address"`
	}{
		Name:    r.Name,
		Email:   r.Email,
		Phone:   r.Phone,
		Address: r.Address,
	}
}

// ToUpdateClientParams converte a request para o formato esperado pelo sqlc
func (r *UpdateClientRequest) ToUpdateClientParams(id int64) interface{} {
	return struct {
		Name    string      `json:"name"`
		Email   string      `json:"email"`
		Phone   pgtype.Text `json:"phone"`
		Address pgtype.Text `json:"address"`
		ID      int64       `json:"id"`
	}{
		Name:    r.Name,
		Email:   r.Email,
		Phone:   r.Phone,
		Address: r.Address,
		ID:      id,
	}
}
