package commentRequest

import (
	"github.com/jackc/pgx/v5/pgtype"

	"sixTask/internal/database"
)

// CreateCommentRequest representa os dados necessários para criar um comentário
// com validações do gin-gonic
type CreateCommentRequest struct {
	Content         string      `json:"content" binding:"required"`
	UserID          pgtype.Int8 `json:"user_id" binding:"required"`
	CommentableType string      `json:"commentable_type" binding:"required"`
	CommentableID   int64       `json:"commentable_id" binding:"required,gt=0"`
}

// UpdateCommentRequest representa os dados necessários para atualizar um comentário
// com validações do gin-gonic
type UpdateCommentRequest struct {
	Content string `json:"content" binding:"required"`
}

// ToCreateCommentParams converte a request para o formato esperado pelo sqlc
func (r *CreateCommentRequest) ToCreateCommentParams() interface{} {
	return database.CreateCommentParams{
		Content:         r.Content,
		UserID:          r.UserID,
		CommentableType: r.CommentableType,
		CommentableID:   r.CommentableID,
	}
}

// ToUpdateCommentParams converte a request para o formato esperado pelo sqlc
func (r *UpdateCommentRequest) ToUpdateCommentParams(id int64) interface{} {
	return database.UpdateCommentParams{
		Content: r.Content,
		ID:      id,
	}
}
