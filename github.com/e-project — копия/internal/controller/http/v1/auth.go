package v1

import (
	"e-project/internal/usecase"
	models "e-project/model"
	"e-project/pkg/logger"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type authRouted struct {
	auth         usecase.Auth
	registration usecase.Registration
	logger       logger.Interface
}

var jwtKey = []byte("new_key")

func newAuthRoutes(handler *gin.RouterGroup, t usecase.Auth, l logger.Interface) {
	r := &authRouted{t, l}
	//TO DO: нужно поменять название, и сделать реализацию авторизации/решистрации
	// h := handler.Group("/translation") // a group of routes with the prefix /translation
	h := handler
	{
		// h.GET("/history", r.history)           //defines the route /translation/history with the HTTP GET method
		h.POST("/login", r.doLogin) //defines the route /translation/do-translate with the HTTP POST method,
		h.POST("/session/check", r.doSessionCheck)
		h.POST("/refresh", r.doRefresh)
	}
}

func (a *authRouted) doLogin(c *gin.Context) {
	var request models.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		a.logger.Error(err, "http - v1 - doLogin")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	if request.Login == "" {
		a.logger.Error("Login is empty", "http - v1 - doLogin")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	if request.Password == "" {
		a.logger.Error("Password is empty", "http - v1 - doLogin")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	user, err := a.auth.GetUser(c, request)
	if err != nil {
		a.logger.Error("Problems with DB", "http - v1 - doLogin")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	if request.Password != user.Password {
		a.logger.Error("Wrong password", "http - v1 - doLogin")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	expirationTime := time.Now().Add(time.Minute * 60) // записываем время

	claims := &models.Claims{
		UUID: user.UUID,
		User: request.Login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		a.logger.Error("Unknown problem", "http - v1 - doLogin")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	session := models.Session{
		UUID:           user.UUID,
		AccessToken:    tokenString,
		ExpirationDate: expirationTime,
		CreatedDate:    time.Now(),
	}

	err = a.auth.Login(c, session)
	if err != nil {
		a.logger.Error("Problems with DB", "http - v1 - doLogin")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})
}

func (a *authRouted) doSessionCheck(c *gin.Context) {

	token, err := c.Cookie("token")
	if err != nil {
		a.logger.Error("Token is not founded", "http - v1 - doSessionCheck")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	claims := &models.Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) { // разбор токена
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			a.logger.Error("Token is wrong", "http - v1 - doSessionCheck")
			errorResponse(c, http.StatusUnauthorized, "invalid request body")

			return
		}
		a.logger.Error("Token is expired", "http - v1 - doSessionCheck")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	if !tkn.Valid {
		a.logger.Error("Token is wrong", "http - v1 - doSessionCheck")
		errorResponse(c, http.StatusUnauthorized, "invalid request body")

		return
	}

	session := models.Session{
		UUID:        claims.UUID,
		AccessToken: token,
	}

	err = a.auth.CheckSessions(c, session)
	if err != nil {
		a.logger.Error("Token expired or not exist", "http - v1 - doSessionCheck")
		errorResponse(c, http.StatusUnauthorized, "invalid request body")

		return
	}

}

func (a *authRouted) doRefresh(c *gin.Context) {
	oldToken, err := c.Cookie("token")
	if err != nil {
		a.logger.Error("Token is not founded", "http - v1 - doRefresh")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	claims := &models.Claims{}

	tkn, err := jwt.ParseWithClaims(oldToken, claims, func(t *jwt.Token) (interface{}, error) { // разбор токена
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			a.logger.Error("Token is wrong", "http - v1 - doRefresh")
			errorResponse(c, http.StatusUnauthorized, "invalid request body")

			return
		}
		a.logger.Error("Token is expired", "http - v1 - doSessionCheck")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	if !tkn.Valid {
		a.logger.Error("Token is wrong", "http - v1 - doRefresh")
		errorResponse(c, http.StatusUnauthorized, "invalid request body")

		return
	}

	session := models.Session{
		UUID:        claims.UUID,
		AccessToken: oldToken,
	}

	err = a.auth.CheckSessions(c, session)
	if err != nil {
		a.logger.Error("Token expired or not exist", "http - v1 - doSessionCheck")
		errorResponse(c, http.StatusUnauthorized, "invalid request body")

		return
	}

	expitarionTime := time.Now().Add(time.Minute * 60)
	claims.ExpiresAt = expitarionTime.Unix()

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := newToken.SignedString(jwtKey)
	if err != nil {
		a.logger.Error("Can't create new token", "http - v1 - doSessionCheck")
		errorResponse(c, http.StatusUnauthorized, "invalid request body")

		return
	}

	session.CreatedDate = time.Now()
	session.AccessToken = tokenString
	session.ExpirationDate = expitarionTime

	err = a.auth.Refresh(c, session, oldToken)
	if err != nil {
		a.logger.Error("Problems with DB", "http - v1 - doSessionCheck")
		errorResponse(c, http.StatusUnauthorized, "invalid request body")

		return
	}

	c.SetCookie("token", tokenString, int(expitarionTime.Unix()), "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Refresh successful",
	})
}
