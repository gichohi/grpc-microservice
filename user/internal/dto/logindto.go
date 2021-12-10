package dto

type LoginDto struct {
	Email 		string    `json:"email" validate:"required,gte=0,lte=255"`
	Password 	string	  `json:"password" validate:"required,gte=0,lte=255"`
}
