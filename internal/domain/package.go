package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type PackageStatus string

const (
	PackageStatusCreated PackageStatus = "created"
	PackageStatusReady   PackageStatus = "ready"
	PackageStatusRunning PackageStatus = "running"
	PackageStatusFailed  PackageStatus = "failed"
)

type ValidationPackage struct {
	UUID            uuid.UUID
	TaxonomyVersion string
	DraftVersion    string
	EntryPointURI   string
	RefPeriodEnd    time.Time
	Status          PackageStatus
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func NewValidationPackage(
	taxonomyVersion string,
	draftVersion string,
	entryPointURI string,
	refPeriodEnd time.Time,
) (*ValidationPackage, error) {
	if taxonomyVersion == "" {
		return nil, fmt.Errorf("taxonomy version is required")
	}
	if draftVersion == "" {
		return nil, fmt.Errorf("draft version is required")
	}
	if entryPointURI == "" {
		return nil, fmt.Errorf("entry point uri is required")
	}
	if refPeriodEnd.IsZero() {
		return nil, fmt.Errorf("ref period end is required")
	}
	now := time.Now().UTC()
	return &ValidationPackage{
		UUID:            uuid.New(),
		TaxonomyVersion: taxonomyVersion,
		DraftVersion:    draftVersion,
		EntryPointURI:   entryPointURI,
		RefPeriodEnd:    refPeriodEnd,
		Status:          PackageStatusCreated,
		CreatedAt:       now,
		UpdatedAt:       now,
	}, nil
}
