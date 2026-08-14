package repository

import "database/sql"

type repo struct {
	db *sql.DB
}

func (r *repo) Save(url *domain.ShortURL) error {
	_, err := r.db.Exec(
		"INSERT INTO urls (id, original) VALUES ($1, $2)",
		url.ID, url.Original,
	)
	return err
}
