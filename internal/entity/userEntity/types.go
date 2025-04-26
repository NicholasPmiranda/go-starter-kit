package userEntity

type UserEntity struct {
	ID    int    `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}
