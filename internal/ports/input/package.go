package input

import (
	"context"

	"github.com/lipkerton/xbrl-validator/internal/domain"
)

type CreatePackageCommand struct {
	TaxonomyVersion string
	DraftVersion    string
	EntryPointURI   string
	RefPeriodEnd    string
}

type PackageUseCase interface {
	CreatePackage(ctx context.Context, cmd CreatePackageCommand) (*domain.ValidationPackage, error)
	GetPackage(ctx context.Context, id string) (*domain.ValidationPackage, error)
}
