package input

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lipkerton/xbrl-validator/internal/domain"
)

type CreatePackageCommand struct {
	TaxonomyVersion string
	DraftVersion    string
	EntryPointURI   string
	RefPeriodEnd    time.Time
}

type PackageUseCase interface {
	CreatePackage(ctx context.Context, cmd CreatePackageCommand) (*domain.ValidationPackage, error)
	GetPackage(ctx context.Context, id uuid.UUID) (*domain.ValidationPackage, error)
}
