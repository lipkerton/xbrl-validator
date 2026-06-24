package domain

type PackageStatus string

const (
	PackageStatusCreated PackageStatus = "created"
	PackageStatusReady   PackageStatus = "ready"
	PackageStatusRunning PackageStatus = "running"
	PackageStatusFailed  PackageStatus = "failed"
)
