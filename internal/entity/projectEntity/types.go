package projectEntity

import "sixTask/internal/entity/userEntity"

type Project struct {
	Name        string                  `json:"name" binding:"required,min=3,max=100"`
	Description string                  `json:"description" binding:"omitempty"`
	ClientID    int                     `json:"client_id" binding:"required"`
	Status      string                  `json:"status" binding:"required"`
	StartDate   string                  `json:"start_date" binding:"omitempty"`
	EndDate     string                  `json:"end_date" binding:"omitempty"`
	UsersId     []userEntity.UserEntity `json:"users" binding:"required"`
}
