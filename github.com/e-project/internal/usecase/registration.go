package usecase

import (
	"context"
	models "e-project/model"
	"fmt"

	"github.com/google/uuid"
)

// TranslationUseCase -.
type RegistrationUseCase struct {
	repo RegistrationRepo
}

// New -.
func New(r RegistrationRepo) *RegistrationUseCase {
	return &RegistrationUseCase{
		repo: r,
	}
}

// History - getting translate history from store.
func (uc *RegistrationUseCase) SaveRegistrationData(ctx context.Context, request models.RegistrationRequest) error {
	id := uuid.New()
	stringId := id.String()
	user := models.User{
		UUID:      stringId,
		FirstName: request.Name,
		SurName:   request.SurName,
		Password:  request.Password,
		Email:     request.Email,
	}

	err := uc.repo.SaveUserData(ctx, user)
	if err != nil {
		return fmt.Errorf("RegistrationUseCase - SaveRegistrationData - uc.repo.SaveUserData: %w", err)
	}

	return nil
}
