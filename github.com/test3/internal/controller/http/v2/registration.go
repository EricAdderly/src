package v2

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"

// 	"e-project/internal/usecase"
// 	models "e-project/model"
// 	"e-project/pkg/logger"
// )

// type registrationRouted struct {
// 	registration usecase.Registration
// 	logger       logger.Interface
// }

// func newRegistrationRoutes(handler *gin.RouterGroup, t usecase.Registration, l logger.Interface) {
// 	r := &registrationRouted{t, l}
// 	h := handler
// 	{
// 		h.POST("/join", r.doRegistration) //defines the route /translation/do-translate with the HTTP POST method,
// 	}
// }

// // @Summary     Translate
// // @Description Translate a text
// // @ID          do-translate
// // @Tags  	    translation
// // @Accept      json
// // @Produce     json
// // @Param       request body doTranslateRequest true "Set up translation"
// // @Success     200 {object} entity.Translation
// // @Failure     400 {object} response
// // @Failure     500 {object} response
// // @Router      /translation/do-translate [post]
// func (r *registrationRouted) doRegistration(c *gin.Context) {
// 	var request models.RegistrationRequest

// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		r.logger.Error(err, "http - v2 - doRegistration")
// 		errorResponse(c, http.StatusBadRequest, "invalid request body")

// 		return
// 	}

// 	if request.Name == "" {
// 		r.logger.Error("Name is empty", "http - v2 - doRegistration")
// 		errorResponse(c, http.StatusBadRequest, "invalid request body")

// 		return
// 	}

// 	if request.SurName == "" {
// 		r.logger.Error("SurName is empty", "http - v2 - doRegistration")
// 		errorResponse(c, http.StatusBadRequest, "invalid request body")

// 		return
// 	}

// 	if request.Email == "" {
// 		r.logger.Error("Email is empty", "http - v2 - doRegistration")
// 		errorResponse(c, http.StatusBadRequest, "invalid request body")

// 		return
// 	}

// 	if request.Password == "" {
// 		r.logger.Error("Password is empty", "http - v2 - doRegistration")
// 		errorResponse(c, http.StatusBadRequest, "invalid request body")

// 		return
// 	}

// 	err := r.registration.SaveRegistrationData(c, request)
// 	if err != nil {
// 		r.logger.Error("Problems with DB", "http - v2 - doRegistration")
// 		errorResponse(c, http.StatusBadRequest, "invalid request body")

// 		return
// 	}
// }
