package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type RegistrationRequest struct {
	Name     string `json:"name"`
	SurName  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	UUID      string
	FirstName string `json:"name"`
	SurName   string `json:"surname"`
	DateBirth string
	Avatar    string
	Location  string
	Role      int
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Claims struct {
	UUID               string `json:"uuid"`
	User               string `json:"login"`
	jwt.StandardClaims `json:"jwt"`
}

type Session struct {
	UUID           string
	AccessToken    string
	ExpirationDate time.Time
	CreatedDate    time.Time
}

// type User struct {
// 	ID       uuid.UUID
// 	Login    string
// 	Email    string
// 	Password string

// 	FirstName   string
// 	SecondName  string
// 	DateOfBirth string
// 	Avatar      []byte
// 	Geolocation string

// 	Verified  bool
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }

// type RegistrationUserInput struct {
// 	Name     string
// 	Email    string
// 	Password string
// }

// type LoginUserInput struct {
// 	Login    string
// 	Password string
// }
