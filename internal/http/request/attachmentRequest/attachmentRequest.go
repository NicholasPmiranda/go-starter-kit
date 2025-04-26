package attachmentRequest

import (
	"github.com/jackc/pgx/v5/pgtype"

	"sixTask/internal/database"
)

// CreateAttachmentRequest representa os dados necessários para criar um anexo
// com validações do gin-gonic
type CreateAttachmentRequest struct {
	Filename       string      `json:"filename" binding:"required"`
	Filepath       string      `json:"filepath" binding:"required"`
	Filesize       int64       `json:"filesize" binding:"required,gt=0"`
	Filetype       string      `json:"filetype" binding:"required"`
	UserID         pgtype.Int8 `json:"user_id" binding:"required"`
	AttachableType string      `json:"attachable_type" binding:"required"`
	AttachableID   int64       `json:"attachable_id" binding:"required,gt=0"`
}

// UpdateAttachmentRequest representa os dados necessários para atualizar um anexo
// com validações do gin-gonic
type UpdateAttachmentRequest struct {
	Filename string `json:"filename" binding:"required"`
	Filepath string `json:"filepath" binding:"required"`
	Filesize int64  `json:"filesize" binding:"required,gt=0"`
	Filetype string `json:"filetype" binding:"required"`
}

// ToCreateAttachmentParams converte a request para o formato esperado pelo sqlc
func (r *CreateAttachmentRequest) ToCreateAttachmentParams() interface{} {
	return database.CreateAttachmentParams{
		Filename:       r.Filename,
		Filepath:       r.Filepath,
		Filesize:       r.Filesize,
		Filetype:       r.Filetype,
		UserID:         r.UserID,
		AttachableType: r.AttachableType,
		AttachableID:   r.AttachableID,
	}
}

// ToUpdateAttachmentParams converte a request para o formato esperado pelo sqlc
func (r *UpdateAttachmentRequest) ToUpdateAttachmentParams(id int64) interface{} {
	return database.UpdateAttachmentParams{
		Filename: r.Filename,
		Filepath: r.Filepath,
		Filesize: r.Filesize,
		Filetype: r.Filetype,
		ID:       id,
	}
}
