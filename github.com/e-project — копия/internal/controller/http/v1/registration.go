package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"e-project/internal/usecase"
	models "e-project/model"
	"e-project/pkg/logger"
)

// TO DO поменять полностью контроллерн на авторизацию и аутентификацию
// type translationRoutes struct {
// 	t usecase.Translation //usecaseAuth
// 	l logger.Interface
// }

type registrationRouted struct {
	registration usecase.Registration
	logger       logger.Interface
}

func newRegistrationRoutes(handler *gin.RouterGroup, t usecase.Registration, l logger.Interface) {
	r := &registrationRouted{t, l}
	//TO DO: нужно поменять название, и сделать реализацию авторизации/решистрации
	// h := handler.Group("/translation") // a group of routes with the prefix /translation
	h := handler
	{
		// h.GET("/history", r.history)           //defines the route /translation/history with the HTTP GET method
		h.POST("/join", r.doRegistration) //defines the route /translation/do-translate with the HTTP POST method,
		// h.POST("/login", r.doLogin)
		// h.POST("/session/check", r.doSessionCheck)
		// h.POST("/refresh", r.doRefresh)
	}
}

// @Summary     Show history
// @Description Show all translation history
// @ID          history
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Success     200 {object} historyResponse
// @Failure     500 {object} response
// @Router      /translation/history [get]
// func (r *translationRoutes) history(c *gin.Context) {
// 	translations, err := r.t.History(c.Request.Context())
// 	if err != nil {
// 		r.l.Error(err, "http - v1 - history")
// 		errorResponse(c, http.StatusInternalServerError, "database problems")

// 		return
// 	}

// 	c.JSON(http.StatusOK, historyResponse{translations})
// }

//	type doTranslateRequest struct {
//		Source      string `json:"source"       binding:"required"  example:"auto"`
//		Destination string `json:"destination"  binding:"required"  example:"en"`
//		Original    string `json:"original"     binding:"required"  example:"текст для перевода"`
//	}

type RegistrationRequest struct {
	Name     string `json:"name"`
	SurName  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @Summary     Translate
// @Description Translate a text
// @ID          do-translate
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Param       request body doTranslateRequest true "Set up translation"
// @Success     200 {object} entity.Translation
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /translation/do-translate [post]
func (r *registrationRouted) doRegistration(c *gin.Context) {
	var request models.RegistrationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.logger.Error(err, "http - v1 - doRegistration")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	if request.Name == "" {
		r.logger.Error("Name is empty", "http - v1 - doRegistration")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	if request.SurName == "" {
		r.logger.Error("SurName is empty", "http - v1 - doRegistration")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	if request.Email == "" {
		r.logger.Error("Email is empty", "http - v1 - doRegistration")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	if request.Password == "" {
		r.logger.Error("Password is empty", "http - v1 - doRegistration")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	id := uuid.New()
	stringId := id.String()
	user := models.User{
		UUID:      stringId,
		FirstName: request.Name,
		SurName:   request.SurName,
		Password:  request.Password,
		Email:     request.Email,
	}
	err := r.registration.SaveRegistrationData(c, user)
	if err != nil {
		r.logger.Error("Problems with DB", "http - v1 - doRegistration")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}
}
