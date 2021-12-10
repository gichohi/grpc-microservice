package dto

type UserDto struct {
	FirstName   string    `json:"firstname" validate:"required,gte=0,lte=255"`
	LastName   string    `json:"lastname" validate:"required,gte=0,lte=255"`
	Email 		string    `json:"email" validate:"required,gte=0,lte=255"`
	Password 	string	  `json:"password" validate:"required,gte=0,lte=255"`
}
