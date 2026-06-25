package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lipkerton/xbrl-validator/internal/domain"
)

type PackageRepository struct {
	db *pgxpool.Pool
}

func NewPackageRepository(db *pgxpool.Pool) *PackageRepository {
	return &PackageRepository{
		db: db,
	}
}

func (r *PackageRepository) Create(ctx context.Context, pkg *domain.ValidationPackage) error {
	query := `
		INSERT INTO validation_packages (
			uuid,
			taxonomy_version,
			draft_version,
			entry_point_url,
			ref_period_end,
			status,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := r.db.Exec(
		ctx,
		query,
		pkg.UUID,
		pkg.TaxonomyVersion,
		pkg.DraftVersion,
		pkg.EntryPointURI,
		pkg.RefPeriodEnd,
		pkg.Status,
		pkg.CreatedAt,
		pkg.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("insert validation package: %w", err)
	}
	return nil
}

func (r *PackageRepository) GetByUUID(ctx context.Context, id uuid.UUID) (*domain.ValidationPackage, error) {
	query := `
		SELECT
			uuid,
			taxonomy_version,
			draft_version,
			entry_point_uri,
			ref_period_end,
			status,
			created_at,
			updated_at
		FROM validation_packages
		WHERE uuid = $1
	`
	var pkg domain.ValidationPackage
	var status string

	err := r.db.QueryRow(ctx, query, id).Scan(
		&pkg.UUID,
		&pkg.TaxonomyVersion,
		&pkg.DraftVersion,
		&pkg.EntryPointURI,
		&pkg.RefPeriodEnd,
		&status,
		&pkg.CreatedAt,
		&pkg.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("package not found: %w", err)
		}

		return nil, fmt.Errorf("select validation package: %w", err)
	}
	pkg.Status = domain.PackageStatus(status)

	return &pkg, nil
}
