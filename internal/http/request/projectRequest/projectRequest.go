package projectRequest

import (
	"github.com/jackc/pgx/v5/pgtype"

	"sixTask/internal/database"
)

// CreateProjectRequest representa os dados necessários para criar um projeto
// com validações do gin-gonic
type CreateProjectRequest struct {
	Name        string        `json:"name" binding:"required,min=3,max=100"`
	Description pgtype.Text   `json:"description" binding:"omitempty"`
	ClientID    pgtype.Int8   `json:"client_id" binding:"required"`
	Status      string        `json:"status" binding:"required"`
	StartDate   pgtype.Date   `json:"start_date" binding:"omitempty"`
	EndDate     pgtype.Date   `json:"end_date" binding:"omitempty"`
	UsersId     []pgtype.Int8 `json:"users_id" binding:"required"`
}

// UpdateProjectRequest representa os dados necessários para atualizar um projeto
// com validações do gin-gonic
type UpdateProjectRequest struct {
	Name        string        `json:"name" binding:"required,min=3,max=100"`
	Description pgtype.Text   `json:"description" binding:"omitempty"`
	ClientID    pgtype.Int8   `json:"client_id" binding:"required"`
	Status      string        `json:"status" binding:"required"`
	StartDate   pgtype.Date   `json:"start_date" binding:"omitempty"`
	EndDate     pgtype.Date   `json:"end_date" binding:"omitempty"`
	UsersId     []pgtype.Int8 `json:"users_id" binding:"required"`
}

// ToCreateProjectParams converte a request para o formato esperado pelo sqlc
func (r *CreateProjectRequest) ToCreateProjectParams() interface{} {
	return database.CreateProjectParams{
		Name:        r.Name,
		Description: r.Description,
		ClientID:    r.ClientID,
		Status:      r.Status,
		StartDate:   r.StartDate,
		EndDate:     r.EndDate,
	}
}

// ToUpdateProjectParams converte a request para o formato esperado pelo sqlc
func (r *UpdateProjectRequest) ToUpdateProjectParams(id int64) interface{} {
	return database.UpdateProjectParams{
		Name:        r.Name,
		Description: r.Description,
		ClientID:    r.ClientID,
		Status:      r.Status,
		StartDate:   r.StartDate,
		EndDate:     r.EndDate,
		ID:          id,
	}
}
