package packageusecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/lipkerton/xbrl-validator/internal/domain"
	"github.com/lipkerton/xbrl-validator/internal/ports/input"
	"github.com/lipkerton/xbrl-validator/internal/ports/output"
)

type Service struct {
	packages output.PackageRepository
}

func NewService(packages output.PackageRepository) *Service {
	return &Service{
		packages: packages,
	}
}

func (s *Service) CreatePackage(ctx context.Context, cmd input.CreatePackageCommand) (*domain.ValidationPackage, error) {
	pkg, err := domain.NewValidationPackage(
		cmd.TaxonomyVersion,
		cmd.DraftVersion,
		cmd.EntryPointURI,
		cmd.RefPeriodEnd,
	)
	if err != nil {
		return nil, fmt.Errorf("build validation package: %w", err)
	}

	if err := s.packages.Create(ctx, pkg); err != nil {
		return nil, fmt.Errorf("save validation package: %w", err)
	}

	return pkg, nil
}

func (s *Service) GetPackage(ctx context.Context, id string) (*domain.ValidationPackage, error) {
	packageUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid package uuid: %w", err)
	}

	pkg, err := s.packages.GetByUUID(ctx, packageUUID)
	if err != nil {
		return nil, fmt.Errorf("get validation package: %w", err)
	}

	return pkg, nil
}
