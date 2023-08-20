CREATE SCHEMA IF NOT EXISTS rinha;

CREATE OR REPLACE FUNCTION extract_names(jsonb) RETURNS TEXT[] AS $$
    SELECT ARRAY(SELECT value->>'name' FROM jsonb_array_elements($1) AS value)
$$ LANGUAGE SQL IMMUTABLE;



CREATE TABLE IF NOT EXISTS rinha.person(
    id VARCHAR(36) NOT NULL,
    nickname VARCHAR(32) PRIMARY KEY NOT NULL,
    name VARCHAR(100) NOT NULL,
    birth_date DATE NOT NULL,
    stack JSONB,    
    full_search TEXT GENERATED ALWAYS AS (
        LOWER(name) || LOWER(nickname) || extract_names(stack)
    ) STORED
);

CREATE EXTENSION pg_trgm;
CREATE INDEX CONCURRENTLY IF NOT EXISTS i_ilikeseach on rinha.person using gist(full_search GIST_TRGM_OPS);


