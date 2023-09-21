package repo

import (
	"context"
	models "e-project/model"
	"e-project/pkg/postgres"
	"fmt"
)

const _defaultEntityCap = 64

// TranslationRepo -.
type RegistrationRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *RegistrationRepo {
	return &RegistrationRepo{pg}
}

func (r *RegistrationRepo) SaveUserData(ctx context.Context, rq models.User) error {
	sql, args, err := r.Builder.
		Insert("public.users").
		Columns("uuid", "first_name, last_name, email, password").
		Values(rq.UUID, rq.FirstName, rq.SurName, rq.Email, rq.Password).
		ToSql()
	if err != nil {
		return fmt.Errorf("Registration - Registrate - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("Registration - Registrate - r.Pool.Exec: %w", err)
	}

	return nil

}
