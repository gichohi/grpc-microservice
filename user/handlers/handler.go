package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/micro/user/internal/models"
	"github.com/micro/user/internal/repository"
	"github.com/micro/user/rpc"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func NewHandler() http.Handler {
	router := chi.NewRouter()
	router.Route("/users", routes)
	return router
}

func routes(router chi.Router) {
	router.Post("/", CreateUser)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	user.Password = string(bytes)

	repository.CreateUser(&user)

	go rpc.Send(&user)

	resp := make(map[string] string)
	resp["userid"] = user.UserID.String()

	msg, err := json.Marshal(resp)
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(msg)
	if err != nil {
		log.Fatalf("Write Error: %s", err)
	}
}
