package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
)

func (r *PostgresStorage) GetFullURL(
	ctx context.Context,
	id int64,
) (string, error) {

	const q = `
	  SELECT full_url
		FROM public.links
		WHERE id = $1
	`

	var fullURL string

	err := r.dbPool.QueryRow(ctx, q, id).Scan(&fullURL)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", errNotFound
		}
		return "", err
	}

	return fullURL, nil
}

var errNotFound = errors.New("The requested shortCode not found in database")
