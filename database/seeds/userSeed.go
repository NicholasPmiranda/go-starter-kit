package seeds

import (
	authhelper "sixTask/helpers/authHelper"
	"sixTask/internal/database"
)

func UserSeeder() {

	passwordHashed, _ := authhelper.HashPassword("admin2024")

	dbCoon, ctx := database.ConnectDB()
	defer dbCoon.Close(ctx)

	query := database.New(dbCoon)

	query.CreateUser(ctx, database.CreateUserParams{
		Name:     "admin",
		Email:    "admin@admin.com",
		Password: passwordHashed,
	})
}
