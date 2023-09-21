package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("new_key")

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type LoginUserInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Claims struct {
	User               string `json:"login"`
	jwt.StandardClaims `json:"jwt"`
}

//	type Session strucr {
//		id uuid.UUID,
//		// погуглить
//	}
func Login(w http.ResponseWriter, r *http.Request) {
	var loginUserInput LoginUserInput
	err := json.NewDecoder(r.Body).Decode(&loginUserInput) // записываем данные из запроса
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPass, ok := users[loginUserInput.Login]

	if !ok || expectedPass != loginUserInput.Password { // разделить на 2
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expitarionTime := time.Now().Add(time.Minute * 60) // записываем время

	claims := &Claims{
		User: loginUserInput.Login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expitarionTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	fmt.Println(err)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expitarionTime,
		})
}

func Home(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token") // записываем куку
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) { // разбор токена
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Hello, %s", claims.User)))
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expitarionTime := time.Now().Add(time.Minute * 60) // записываем текущее время

	claims.ExpiresAt = expitarionTime.Unix() // перезаписываем его в токен

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "refresh_token",
			Value:   tokenString,
			Expires: expitarionTime,
		})
}

type RegistrationRequest struct {
	Name       string `json:"name"`
	SecondName string `json:"secon_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	err := RegistrationController(w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func RegistrationController(w http.ResponseWriter, r *http.Request) error {
	var reg RegistrationRequest

	err := json.NewDecoder(r.Body).Decode(&reg)
	if err != nil {
		return err
	}

	if reg.Name == "" || reg.SecondName == "" || reg.Email == "" || reg.Password == "" {
		return err
	}

	err = RegistrtionRepository(reg)
	if err != nil {
		return err
	}

	err = SendConfirmationEmail(reg.Email)
	if err != nil {
		return err
	}
	return nil
}

func RegistrtionRepository(reg RegistrationRequest) error {
	var err error
	return err
}

func SendConfirmationEmail(email string) error {
	// Настройки SMTP-сервера который будет отправлять письмо
	smtpHost := "smtp.example.com"
	smtpPort := 587
	smtpUsername := "your-smtp-username"
	smtpPassword := "your-smtp-password"

	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	// Формирование письма
	from := "noreply@example.com"
	to := []string{email}           // список адресов
	subject := "Confirmation Email" // тема письма
	body := "Dear User, please confirm your registration."

	msg := []byte("To: " + email + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body)

	// Отправка письма через SMTP-сервер
	err := smtp.SendMail(smtpHost+":"+string(smtpPort), auth, from, to, msg)
	if err != nil {
		return err
	}

	return nil
}
