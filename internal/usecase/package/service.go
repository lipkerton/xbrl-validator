package packageusecase

import (
	"context"

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

}
