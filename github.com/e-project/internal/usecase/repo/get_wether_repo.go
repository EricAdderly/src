package repo

import "e-project/pkg/postgres"

type GetWeatherRepo struct {
	*postgres.Postgres
}

func NewGetWeatherRepo(pg *postgres.Postgres) *GetWeatherRepo {
	return &GetWeatherRepo{pg}
}
