CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS canvas
(
    id         UUID    PRIMARY KEY DEFAULT uuid_generate_v4(),
    drawing    TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

