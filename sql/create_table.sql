CREATE SCHEMA IF NOT EXISTS rinha;


CREATE TABLE IF NOT EXISTS rinha.person(
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    nickname VARCHAR(32) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    birth_date DATE NOT NULL,
    stack JSONB,
    full_search TSVECTOR
);

CREATE INDEX IF NOT EXISTS i_fulltext on rinha.person using gin(full_search);


CREATE OR REPLACE FUNCTION create_full_text_search() RETURNS TRIGGER AS $$
BEGIN    
    NEW.full_search :=
        TO_TSVECTOR('english', CONCAT_WS(' || ', NEW.nickname, NEW.name,             
            jsonb_path_query_array(New.stack, 'strict $.**.name')
        ));           
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_tsvector_trigger
BEFORE INSERT OR UPDATE ON rinha.person
FOR EACH ROW
EXECUTE FUNCTION create_full_text_search();
