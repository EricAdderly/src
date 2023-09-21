package repo

import (
	"context"
	models "e-project/model"
	"e-project/pkg/postgres"
	"fmt"
)

// TranslationRepo -.
type AuthRepo struct {
	*postgres.Postgres
}

func NewAuthRepo(pg *postgres.Postgres) *AuthRepo {
	return &AuthRepo{pg}
}

func (a *AuthRepo) SaveLoginData(ctx context.Context, session models.Session) error {
	sql, args, err := a.Builder.
		Insert("public.auth_tokens").
		Columns("uuid", "access_token", "expiration_date", "created_date").
		Values(session.UUID, session.AccessToken, session.ExpirationDate, session.CreatedDate).
		ToSql()
	if err != nil {
		return fmt.Errorf("Authrisation - SaveLoginData - a.Builder: %w", err)
	}

	_, err = a.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("Authrisation - SaveLoginData - a.Pool.Exec: %w", err)
	}

	return nil
}

func (a *AuthRepo) CheckSessionsData(ctx context.Context, session models.Session) (models.Session, error) {
	sql, _, err := a.Builder.
		Select("uuid", "expiration_date").
		From("public.auth_tokens").
		Where("access_token = ?", session.AccessToken).
		ToSql()
	if err != nil {
		return models.Session{}, fmt.Errorf("Authrisation - CheckSessionsData - a.Builder: %w", err)
	}

	rows, err := a.Pool.Query(ctx, sql, session.AccessToken)
	if err != nil {
		return models.Session{}, fmt.Errorf("Authrisation - CheckSessionsData - a.Pool.Query: %w", err)
	}
	defer rows.Close()

	sessionDb := models.Session{}

	for rows.Next() {

		err = rows.Scan(&sessionDb.UUID, &sessionDb.ExpirationDate)
		if err != nil {
			return models.Session{}, fmt.Errorf("Authrisation - CheckSessionsData - rows.Scan: %w", err)
		}

	}

	return sessionDb, nil
}

func (a *AuthRepo) GetUserFromDb(ctx context.Context, rq models.LoginRequest) (models.User, error) {
	sql, _, err := a.Builder.
		Select("uuid", "first_name, last_name, email, password").
		From("public.users").
		Where("email = ?", rq.Login).
		ToSql()
	if err != nil {
		return models.User{}, fmt.Errorf("Authrisation - GetUserFromDb - a.Builder: %w", err)
	}

	rows, err := a.Pool.Query(ctx, sql, rq.Login)
	if err != nil {
		return models.User{}, fmt.Errorf("Authrisation - GetUserFromDb - a.Pool.Query: %w", err)
	}
	defer rows.Close()

	user := models.User{}
	for rows.Next() {

		err = rows.Scan(&user.UUID, &user.FirstName, &user.SurName, &user.Email, &user.Password)
		if err != nil {
			return models.User{}, fmt.Errorf("Authrisation - GetUserFromDb - rows.Scan: %w", err)
		}

	}

	return user, nil
}

func (a *AuthRepo) RefeteshSessionDate(ctx context.Context, session models.Session, oldToken string) error {
	sql, args, err := a.Builder.
		Update("public.auth_tokens").
		Set("access_token", session.AccessToken).
		Set("expiration_date", session.ExpirationDate).
		Set("created_date", session.CreatedDate).
		Where("access_token = ?", oldToken).
		ToSql()
	if err != nil {
		return fmt.Errorf("Authrisation - RefeteshSessionDate - a.Builder: %w", err)
	}
	fmt.Println(sql)
	rows, err := a.Pool.Query(ctx, sql, args...)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Authrisation - RefeteshSessionDate - a.Pool.Query: %w", err)
	}
	defer rows.Close()

	return nil
}
