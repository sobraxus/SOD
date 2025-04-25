package repositories

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/sobraxus/SOD/internal/models"
)

type CaseRepository struct {
	db *pgx.Conn
}

func NewCaseRepository(db *pgx.Conn) *CaseRepository {
	return &CaseRepository{db: db}
}

func (r *CaseRepository) CreateCase(ctx context.Context, c *models.Case) error {
	query := `
		INSERT INTO cases (id, title, description, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.Exec(
		ctx,
		query,
		c.ID,
		c.Title,
		c.Description,
		c.Status,
		c.CreatedAt,
		c.UpdatedAt,
	)
	return err
}

func (r *CaseRepository) GetCaseByID(ctx context.Context, id uuid.UUID) (*models.Case, error) {
	query := `
		SELECT id, title, description, status, created_at, updated_at
		FROM cases
		WHERE id = $1
	`
	var c models.Case
	err := r.db.QueryRow(ctx, query, id).Scan(
		&c.ID,
		&c.Title,
		&c.Description,
		&c.Status,
		&c.CreatedAt,
		&c.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get case: %w", err)
	}
	return &c, nil
}
