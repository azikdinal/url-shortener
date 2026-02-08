package postgres

import (
	"context"
	"fmt"
)

// Сохранениние fullURL в памяти для последующего update
func (r *PostgresStorage) SaveFullURL(ctx context.Context, fullURL string) (int64, error) {
	var id int64

	const q = `
    INSERT INTO public.links (full_url)
		VALUES ($1)
		RETURNING id
	`

	err := r.dbPool.QueryRow(ctx, q, fullURL).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, nil
}
