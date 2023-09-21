// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"
	models "e-project/model"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Translation -.
	// Translation interface {
	// 	Translate(context.Context, entity.Translation) (entity.Translation, error)
	// 	History(context.Context) ([]entity.Translation, error)
	// }

	// TranslationRepo -.
	Registration interface {
		SaveRegistrationData(context.Context, models.User) error
	}

	// TranslationWebAPI -.
	// TranslationWebAPI interface {
	// 	Translate(entity.Translation) (entity.Translation, error)
	// }

	//Registration -.
	RegistrationRepo interface {
		SaveUserData(context.Context, models.User) error
	}

	//AUTH -.
	Auth interface {
		Login(context.Context, models.Session) error
		CheckSessions(context.Context, models.Session) error
		Refresh(context.Context, models.Session, string) error
		GetUser(context.Context, models.LoginRequest) (models.User, error)
	}

	AuthRepo interface {
		SaveLoginData(context.Context, models.Session) error
		CheckSessionsData(context.Context, models.Session) (models.Session, error)
		GetUserFromDb(context.Context, models.LoginRequest) (models.User, error)
		RefeteshSessionDate(context.Context, models.Session, string) error
	}
)
