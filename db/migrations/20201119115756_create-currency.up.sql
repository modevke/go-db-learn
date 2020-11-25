CREATE TABLE IF NOT EXISTS currencies (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name STRING NOT NULL,
    code STRING NOT NULL UNIQUE,
    symbol STRING NOT NULL,
    status STRING NOT NULL DEFAULT 'ACTIVE',
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP():::TIMESTAMPTZ,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP():::TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    CONSTRAINT "primary" PRIMARY KEY (id ASC)
);