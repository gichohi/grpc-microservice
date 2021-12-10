package rpc

import (
	"context"
	"github.com/micro/user/email"
	"github.com/micro/user/internal/models"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"log"
	"time"
)

func Send(user *models.User) {
	conn, err := grpc.Dial("todoworld.servicestack.net:5054", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := email.NewEmailServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	
	var email email.Email
	email.Subject = "Welcome"
	email.Uuid = uuid.NewV4().String()
	email.Body = "Welcome " + user.FirstName + " " + user.LastName
	email.Address = user.Email

	send, err := client.Send(ctx, &email)
	if err != nil {
		return 
	}

	log.Printf(send.Uuid)
}
