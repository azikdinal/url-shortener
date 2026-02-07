package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
)

func (r *PostgresStorage) GetFullURL(
	ctx context.Context,
	shortCode string,
) (string, error) {

	const q = `
	  SELECT full_url
		FROM links
		WHERE short_code = $1
	`

	var fullURL string

	err := r.dbPool.QueryRow(ctx, q, shortCode).Scan(&fullURL)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", errNotFound
		}
		return "", err
	}

	return shortCode, nil
}

var errNotFound = errors.New("The requested shortCode not found in database")
