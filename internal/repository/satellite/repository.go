package satellite

import (
	"context"
	"database/sql"

	model "github.com/gritsulyak/awesomeProject/internal/model"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, s model.Satellite) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO satellite (name) VALUES ($1)", s.Name)
	return err
}

func (r *repository) GetByName(ctx context.Context, name string) (*model.Satellite, error) {
	var s model.Satellite
	err := r.db.QueryRowContext(ctx, "SELECT name FROM satellite WHERE name = $1", name).Scan(&s.Name)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
