package repository

import (
	"fmt"
	"github.com/micro/user/internal/db"
	"github.com/micro/user/internal/models"
)

func CreateUser(user *models.User)  string {
	database := db.Connect()

	fmt.Println("User")
	fmt.Println(user)

	sqlStatement := `
	INSERT INTO users (user_id, firstname, lastname, email, password)
	VALUES ($1, $2, $3, $4)
	`
	_, err := database.Exec(sqlStatement, &user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.Password)

	if err != nil {
		fmt.Println("DB Error: ", err)
	}
	return "Created"
}