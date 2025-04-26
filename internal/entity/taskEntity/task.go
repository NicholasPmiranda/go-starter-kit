package taskEntity

import (
	"sixTask/internal/database"
	"sixTask/internal/entity/userEntity"

	"github.com/jackc/pgx/v5/pgtype"
)

// Task representa uma tarefa com informações do usuário associado
type Task struct {
	ID          int64            `json:"id"`
	Title       string           `json:"title"`
	Description pgtype.Text      `json:"description"`
	ProjectID   pgtype.Int8      `json:"project_id"`
	AssignedTo  pgtype.Int8      `json:"assigned_to"`
	Status      string           `json:"status"`
	Priority    string           `json:"priority"`
	DueDate     pgtype.Date      `json:"due_date"`
	CompletedAt pgtype.Timestamp `json:"completed_at"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
	User        *userEntity.User `json:"user,omitempty"`
}

// TaskWithPagination contém as tarefas paginadas com informações de usuário e metadados de paginação
type TaskWithPagination struct {
	Data []Task         `json:"data"`
	Meta PaginationMeta `json:"meta"`
}

// PaginationMeta contém os metadados de paginação
type PaginationMeta struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
	TotalPages  int   `json:"total_pages"`
}

// FromDatabaseTask converte um database.Task para taskEntity.Task
func FromDatabaseTask(dbTask database.Task) Task {
	return Task{
		ID:          dbTask.ID,
		Title:       dbTask.Title,
		Description: dbTask.Description,
		ProjectID:   dbTask.ProjectID,
		AssignedTo:  dbTask.AssignedTo,
		Status:      dbTask.Status,
		Priority:    dbTask.Priority,
		DueDate:     dbTask.DueDate,
		CompletedAt: dbTask.CompletedAt,
		CreatedAt:   dbTask.CreatedAt,
		UpdatedAt:   dbTask.UpdatedAt,
	}
}

// ParseTasksWithUsers combina tarefas e usuários em uma única estrutura
func ParseTasksWithUsers(tasks []database.Task, users []database.User) []Task {
	// Criar um mapa de usuários por ID para facilitar a busca
	userMap := make(map[int64]database.User)
	for _, user := range users {
		userMap[user.ID] = user
	}

	// Combinar tarefas com seus usuários
	result := make([]Task, len(tasks))
	for i, task := range tasks {
		taskEntity := FromDatabaseTask(task)

		// Adicionar informações do usuário se existir
		if task.AssignedTo.Valid {
			if user, exists := userMap[task.AssignedTo.Int64]; exists {
				userEntity := userEntity.FromDatabaseUser(user)
				taskEntity.User = &userEntity
			}
		}

		result[i] = taskEntity
	}

	return result
}
