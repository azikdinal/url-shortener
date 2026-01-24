package repo

import (
	"context"
	"database/sql"
	"errors"
	"shorten/internal/domain"
)

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(db *sql.DB) Repo {
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) GetByCode(
	ctx context.Context,
	sc domain.ShortCode,
) (domain.FullURL, error) {

	codeStr := string(sc)
	row := r.db.QueryRowContext(
		ctx,
		"SELECT full_url FROM links WHERE short_code = $1",
		codeStr,
	)

	var fullURL string
	if err := row.Scan(&fullURL); err != nil {
		return "", err
	}

	fu, err := domain.NewFullURL(fullURL)
	if err != nil {
		return "", errors.New("invalid fullURL found in Postgres")
	}

	return fu, nil
}

func (r *PostgresRepo) Save(ctx context.Context, link *domain.Link) error {
	fu := string(link.FullURL())
	sc := string(link.ShortCode())

	_, err := r.db.ExecContext(ctx,
		`INSERT INTO links (short_code, full_url)
		 VALUES ($1, $2)`,
		sc, fu,
	)
	if err != nil {
		return err
	}

	return nil
}
