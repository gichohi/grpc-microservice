package repository

import (
	"fmt"
	"github.com/micro/user/internal/db"
	"github.com/micro/user/internal/models"
	"log"
)

func CreateUser(user *models.User)  string {
	database := db.Connect()

	sqlStatement := `
	INSERT INTO users (user_id, firstname, lastname, email, password, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, now(), now())
	`
	_, err := database.Exec(sqlStatement, &user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.Password)

	if err != nil {
		fmt.Println("DB Error: ", err)
	}
	return "Created"
}

func GetUser(email string) *models.User {
	database := db.Connect()

	var user models.User

	sqlStatement := `
     SELECT u.user_id, u.firstname, u.lastname, u.email,u.password, u.created_at, u.updated_at 
	FROM users u WHERE u.email = $1
     `
	err := database.QueryRow(sqlStatement, email).Scan(
		&user.UserID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		)
	if err != nil {
		log.Printf("Get Error: ", err)
	}
	return &user
}