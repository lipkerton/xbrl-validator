package output

import (
	"context"

	"github.com/google/uuid"
	"github.com/lipkerton/xbrl-validator/internal/domain"
)

type PackageRepository interface {
	Create(ctx context.Context, pkg *domain.ValidationPackage) error
	GetByUUID(ctx context.Context, id uuid.UUID) (*domain.ValidationPackage, error)
}
