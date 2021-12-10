package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/micro/user/internal/dto"
	"github.com/micro/user/internal/models"
	"github.com/micro/user/internal/repository"
	"github.com/micro/user/rpc"
	"github.com/micro/user/util"
	uuid "github.com/satori/go.uuid"
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
	router.Post("/login", LoginUser)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var userdto dto.UserDto

	err := json.NewDecoder(r.Body).Decode(&userdto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	password := util.HashPassword(userdto.Password)

	var user models.User
	user.UserID = uuid.NewV4()
	user.Email = userdto.Email
	user.FirstName = userdto.FirstName
	user.LastName = userdto.LastName
	user.Password = password

	res := repository.CreateUser(&user)

	fmt.Println(res)

	go rpc.Send(&user)

	resp := make(map[string] string)
	resp["user_id"] = user.UserID.String()

	msg, err := json.Marshal(resp)
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(msg)
	if err != nil {
		log.Fatalf("Write Error: %s", err)
	}
}


func LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginDto dto.LoginDto

	err := json.NewDecoder(r.Body).Decode(&loginDto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	email := loginDto.Email
	pwd := loginDto.Password

	user := repository.GetUser(email)
	password := user.Password


	if comparePasswords(password, pwd) {
		token := util.GenerateToken(email)
		msg, err := json.Marshal(dto.ResponseDto{Token: token})
		w.WriteHeader(http.StatusCreated)
		_, err = w.Write(msg)
		if err != nil {
			log.Fatalf("Write Error: %s", err)
		}
	} else {
		msg, err := json.Marshal(dto.ErrorDto{Error: "Incorrect Email/Password"})
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write(msg)
		if err != nil {
			log.Fatalf("Write Error: %s", err)
		}
	}
}

func comparePasswords(hashedPwd string, pwd string) bool {
	byteHash := []byte(hashedPwd)
	plainPwd := []byte(pwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println("Compare: " , err)
		return false
	}

	return true
}
