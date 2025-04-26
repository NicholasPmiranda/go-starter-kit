package userEntity

import "sixTask/internal/database"

// User representa um usuário no sistema
type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// FromDatabaseUser converte um database.User para userEntity.User
func FromDatabaseUser(dbUser database.User) User {
	return User{
		ID:    dbUser.ID,
		Name:  dbUser.Name,
		Email: dbUser.Email,
	}
}
