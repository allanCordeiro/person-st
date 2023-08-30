CREATE TABLE IF NOT EXISTS person(
    id VARCHAR(36) UNIQUE NOT NULL,
    nickname VARCHAR(32) PRIMARY KEY NOT NULL,
    name VARCHAR(100) NOT NULL,
    birth_date DATE NOT NULL,
    stack JSONB,    
    full_search TEXT
);

CREATE EXTENSION pg_trgm;
CREATE INDEX CONCURRENTLY IF NOT EXISTS i_ilikeseach on person using gist(full_search GIST_TRGM_OPS);
