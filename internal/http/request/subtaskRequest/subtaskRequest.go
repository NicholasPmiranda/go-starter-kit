package subtaskRequest

import (
	"github.com/jackc/pgx/v5/pgtype"

	"sixTask/internal/database"
)

// CreateSubtaskRequest representa os dados necessários para criar uma subtarefa
// com validações do gin-gonic
type CreateSubtaskRequest struct {
	Title       string      `json:"title" binding:"required,min=3,max=100"`
	Description pgtype.Text `json:"description" binding:"omitempty"`
	TaskID      pgtype.Int8 `json:"task_id" binding:"required"`
	AssignedTo  pgtype.Int8 `json:"assigned_to" binding:"required"`
	Status      string      `json:"status" binding:"required"`
	DueDate     pgtype.Date `json:"due_date" binding:"omitempty"`
}

// UpdateSubtaskRequest representa os dados necessários para atualizar uma subtarefa
// com validações do gin-gonic
type UpdateSubtaskRequest struct {
	Title       string      `json:"title" binding:"required,min=3,max=100"`
	Description pgtype.Text `json:"description" binding:"omitempty"`
	TaskID      pgtype.Int8 `json:"task_id" binding:"required"`
	AssignedTo  pgtype.Int8 `json:"assigned_to" binding:"required"`
	Status      string      `json:"status" binding:"required"`
	DueDate     pgtype.Date `json:"due_date" binding:"omitempty"`
}

// ToCreateSubtaskParams converte a request para o formato esperado pelo sqlc
func (r *CreateSubtaskRequest) ToCreateSubtaskParams() interface{} {
	return database.CreateSubtaskParams{
		Title:       r.Title,
		Description: r.Description,
		TaskID:      r.TaskID,
		AssignedTo:  r.AssignedTo,
		Status:      r.Status,
		DueDate:     r.DueDate,
	}
}

// ToUpdateSubtaskParams converte a request para o formato esperado pelo sqlc
func (r *UpdateSubtaskRequest) ToUpdateSubtaskParams(id int64) interface{} {
	return database.UpdateSubtaskParams{
		Title:       r.Title,
		Description: r.Description,
		TaskID:      r.TaskID,
		AssignedTo:  r.AssignedTo,
		Status:      r.Status,
		DueDate:     r.DueDate,
		ID:          id,
	}
}
