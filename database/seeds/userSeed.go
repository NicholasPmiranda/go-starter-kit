package seeds

import (
	authhelper "boilerPlate/helpers/authHelper"
	"boilerPlate/internal/database"
)

func UserSeeder() {

	passwordHashed, _ := authhelper.HashPassword("admin")

	dbCoon, ctx := database.ConnectDB()
	defer dbCoon.Close(ctx)

	query := database.New(dbCoon)

	query.CreateUser(ctx, database.CreateUserParams{
		Name:     "admin",
		Email:    "amdin@amdin.com",
		Password: passwordHashed,
	})
}
