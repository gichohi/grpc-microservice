package email

import "context"

type Server struct {
}

func (s *Server) Send(ctx context.Context,e *Email) (*Email, error) {
	return &Email{Uuid: e.Uuid, Subject: e.Subject, Address: e.Address , Body: e.Body}, nil
}