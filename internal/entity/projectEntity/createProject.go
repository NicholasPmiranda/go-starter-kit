package projectEntity

import (
	"sixTask/internal/database"
	"sixTask/internal/entity/userEntity"
)

func GetProjectEntity(project database.Project, users []database.User) Project {
	// Create a slice to hold user responses
	userResponses := make([]userEntity.User, 0, len(users))

	// Convert each database.User to UserResponse
	for _, user := range users {
		userResponses = append(userResponses, userEntity.User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	// Create and return the Project
	return Project{
		Name:        project.Name,
		Description: project.Description.String,
		ClientID:    int(project.ClientID.Int64),
		Status:      project.Status,
		StartDate:   project.StartDate.Time.Format("2006-01-02"),
		EndDate:     project.EndDate.Time.Format("2006-01-02"),
		UsersId:     userResponses,
	}
}
