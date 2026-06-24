-- +goose Up
CREATE TABLE validation_packages (
    uuid UUID PRIMARY KEY,
    taxonomy_version TEXT NOT NULL,
    draft_version TEXT NOT NULL,
    entry_point_uri TEXT NOT NULL,
    ref_period_end DATE NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE uploaded_files (
    id UUID PRIMARY KEY,
    package_uuid UUID NOT NULL REFERENCES validation_packages(uuid) ON DELETE CASCADE,
    original_name TEXT NOT NULL,
    role_name TEXT,
    stage_table TEXT,
    rows_loaded INTEGER NOT NULL DEFAULT 0,
    mode TEXT NOT NULL,
    status TEXT NOT NULL,
    error_message TEXT,
    upload_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE validation_runs (
    id UUID PRIMARY KEY,
    package_uuid UUID NOT NULL REFERENCES validation_packages(uuid) ON DELETE CASCADE,
    status TEXT NOT NULL,
    formulas_total INTEGER NOT NULL DEFAULT 0,
    formulas_executed INTEGER NOT NULL DEFAULT 0,
    true_results INTEGER NOT NULL DEFAULT 0,
    false_results INTEGER NOT NULL DEFAULT 0,
    result_xlsx_path TEXT,
    error_message TEXT,
    started_at TIMESTAMPTZ,
    finished_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE jobs (
    id UUID PRIMARY KEY,
    type TEXT NOT NULL,
    status TEXT NOT NULL,
    payload JSONB NOT NULL,
    attempts INTEGER NOT NULL DEFAULT 0,
    max_attempts INTEGER NOT NULL DEFAULT 3,
    error_message TEXT,
    available_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    locked_at TIMESTAMPTZ,
    locked_by TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_uploaded_files_package_uuid ON uploaded_files(package_uuid);
CREATE INDEX idx_validation_runs_package_uuid ON validation_runs(package_uuid);
CREATE INDEX idx_jobs_status_available_at ON jobs(status, available_at);

-- +goose Down
DROP TABLE IF EXISTS jobs;
DROP TABLE IF EXISTS validation_runs;
DROP TABLE IF EXISTS uploaded_files;
DROP TABLE IF EXISTS validation_packages;

