package v1

import (
	"e-project/internal/usecase"
	models "e-project/model"
	"e-project/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authRouted struct {
	auth   usecase.Auth
	logger logger.Interface
}

var jwtKey = []byte("new_key")

func newAuthRoutes(handler *gin.RouterGroup, t usecase.Auth, l logger.Interface) {
	r := &authRouted{t, l}
	h := handler
	{
		h.POST("/login", r.doLogin)
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

	session, err := a.auth.Login(c, request, user)
	if err != nil {
		a.logger.Error("Problems with DB", "http - v1 - doLogin")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	c.SetCookie("token", session.AccessToken, int(session.ExpirationDate.Unix()), "/", "", false, true)

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

	err = a.auth.CheckSessions(c, token)
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

	session, err := a.auth.Refresh(c, oldToken)
	if err != nil {
		a.logger.Error("Problems with DB", "http - v1 - doSessionCheck")
		errorResponse(c, http.StatusUnauthorized, "invalid request body")

		return
	}

	c.SetCookie("token", session.AccessToken, int(session.ExpirationDate.Unix()), "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Refresh successful",
	})
}
