package usecase

import (
	"context"
	models "e-project/model"
	"fmt"
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
func (uc *RegistrationUseCase) SaveRegistrationData(ctx context.Context, rq models.User) error {
	err := uc.repo.SaveUserData(ctx, rq)
	if err != nil {
		return fmt.Errorf("RegistrationUseCase - SaveRegistrationData - uc.repo.SaveUserData: %w", err)
	}

	return nil
}

// // Translate -.
// func (uc *RegistrationUseCase) Translate(ctx context.Context, t entity.Translation) (entity.Translation, error) {
// 	translation, err := uc.webAPI.Translate(t)
// 	if err != nil {
// 		return entity.Translation{}, fmt.Errorf("TranslationUseCase - Translate - s.webAPI.Translate: %w", err)
// 	}

// 	err = uc.repo.Store(context.Background(), translation)
// 	if err != nil {
// 		return entity.Translation{}, fmt.Errorf("TranslationUseCase - Translate - s.repo.Store: %w", err)
// 	}

// 	return translation, nil
// }
