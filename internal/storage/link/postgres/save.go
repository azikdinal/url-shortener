package postgres

import (
	"context"
)

func (r *PostgresStorage) Save(ctx context.Context, shortCode string, fullURL string) error {
	const q = `
    INSERT INTO links (short_code, full_url)
		VALUES ($1, $2)
	`

	_, err := r.dbPool.Exec(
		ctx,
		q,
		shortCode,
		fullURL,
	)

	return err
}
