
CREATE TABLE IF NOT EXISTS satellite (
    id   BIGSERIAL    PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);

INSERT INTO satellite (name) VALUES
    ('moon'), ('europa'), ('titan'), ('ganymede')
ON CONFLICT (name) DO NOTHING;