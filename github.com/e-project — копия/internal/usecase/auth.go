package usecase

import (
	"context"
	models "e-project/model"
	"fmt"
	"time"
)

type AuthUseCase struct {
	repo AuthRepo
}

func NewAuth(r AuthRepo) *AuthUseCase {
	return &AuthUseCase{
		repo: r,
	}
}

func (uc *AuthUseCase) Login(c context.Context, session models.Session) error {
	err := uc.repo.SaveLoginData(c, session)
	if err != nil {
		return fmt.Errorf("AuthUseCase - Login - uc.repo.SaveLoginData: %w", err)
	}

	return nil
}

func (uc *AuthUseCase) CheckSessions(ctx context.Context, session models.Session) error {
	rqUUID := session.UUID
	sessionDb, err := uc.repo.CheckSessionsData(ctx, session)
	if err != nil {
		return fmt.Errorf("AuthUseCase - CheckSessions - uc.repo.CheckSessions: %w", err)
	}

	now := time.Now()
	if now.After(sessionDb.ExpirationDate) {
		return fmt.Errorf("AuthUseCase - CheckSessions - now.After(expirationtime: %w", err)
	}

	if rqUUID != sessionDb.UUID {
		return fmt.Errorf("AuthUseCase - CheckSessions - rqUUID != sessionDb.UUID: %w", err)
	}

	return nil
}

func (uc *AuthUseCase) Refresh(ctx context.Context, session models.Session, oldToken string) error {
	err := uc.repo.RefeteshSessionDate(ctx, session, oldToken)
	if err != nil {
		return fmt.Errorf("AuthUseCase - Refresh - uc.repo.RefeteshSessionDate: %w", err)
	}
	return nil
}

func (uc *AuthUseCase) GetUser(ctx context.Context, rq models.LoginRequest) (models.User, error) {
	user, err := uc.repo.GetUserFromDb(ctx, rq)
	if err != nil {
		return models.User{}, fmt.Errorf("AuthUseCase - History - uc.repo.GetHistory: %w", err)
	}

	return user, nil
}
