package usecase

import (
	"context"
	models "e-project/model"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("new_key")

type AuthUseCase struct {
	repo AuthRepo
}

func NewAuth(r AuthRepo) *AuthUseCase {
	return &AuthUseCase{
		repo: r,
	}
}

func (uc *AuthUseCase) Login(ctx context.Context, request models.LoginRequest, user *models.User) (*models.Session, error) {
	session, err := uc.createSession(ctx, request, user)
	if err != nil {
		return nil, fmt.Errorf("AuthUseCase - Login - uc.repo.createSession: %w", err)
	}

	err = uc.repo.SaveLoginData(ctx, session)
	if err != nil {
		return nil, fmt.Errorf("AuthUseCase - Login - uc.repo.SaveLoginData: %w", err)
	}

	return session, nil
}

func (uc *AuthUseCase) createSession(c context.Context, request models.LoginRequest, user *models.User) (*models.Session, error) {
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
		return nil, fmt.Errorf("AuthUseCase - createSession - token.SignedString(jwtKey): %w", err)
	}

	session := models.Session{
		UUID:           user.UUID,
		AccessToken:    tokenString,
		ExpirationDate: expirationTime,
		CreatedDate:    time.Now(),
	}

	return &session, nil
}

func (uc *AuthUseCase) CheckSessions(ctx context.Context, token string) error {
	session, err := uc.checkSession(ctx, token)
	if err != nil {
		return fmt.Errorf("AuthUseCase - CheckSessions - uc.checkSession: %w", err)
	}

	rqUUID := session.UUID

	sessionDb, err := uc.repo.CheckSessionsData(ctx, *session)
	if err != nil {
		return fmt.Errorf("AuthUseCase - CheckSessions - uc.repo.CheckSessionsData: %w", err)
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

func (uc *AuthUseCase) checkSession(ctx context.Context, token string) (*models.Session, error) {
	claims := &models.Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) { // разбор токена
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, fmt.Errorf("AuthUseCase - checkSession - err == jwt.ErrSignatureInvalid: %w", err)

		}
		return nil, fmt.Errorf("AuthUseCase - checkSession - err == jwt.ParseWithClaims: %w", err)
	}

	if !tkn.Valid {
		return nil, fmt.Errorf("AuthUseCase - checkSession - !tkn.Valid: %w", err)
	}

	session := models.Session{
		UUID:        claims.UUID,
		AccessToken: token,
	}

	return &session, err
}

func (uc *AuthUseCase) Refresh(ctx context.Context, oldToken string) (*models.Session, error) {
	session, err := uc.checkSession(ctx, oldToken)
	if err != nil {
		return nil, fmt.Errorf("AuthUseCase - Refresh - uc.checkSession: %w", err)
	}

	claims := &models.Claims{}
	expitarionTime := time.Now().Add(time.Minute * 60)
	claims.ExpiresAt = expitarionTime.Unix()

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := newToken.SignedString(jwtKey)
	if err != nil {
		return nil, fmt.Errorf("AuthUseCase - Refresh - jwt.NewWithClaims: %w", err)
	}

	session.CreatedDate = time.Now()
	session.AccessToken = tokenString
	session.ExpirationDate = expitarionTime

	err = uc.repo.RefeteshSessionDate(ctx, session, oldToken)
	if err != nil {
		return nil, fmt.Errorf("AuthUseCase - Refresh - uc.repo.RefeteshSessionDate: %w", err)
	}

	return session, nil
}

func (uc *AuthUseCase) GetUser(ctx context.Context, rq models.LoginRequest) (*models.User, error) {
	user, err := uc.repo.GetUserFromDb(ctx, rq)
	if err != nil {
		return nil, fmt.Errorf("AuthUseCase - GetUser - uc.repo.GetUserFromDb: %w", err)
	}

	return user, nil
}
