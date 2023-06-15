CREATE TYPE prefix AS ENUM ('Mr', 'Mrs', 'Ms', 'Dr', 'Prof');
CREATE TYPE suffix AS ENUM ('Jr', 'Sr', 'II', 'III', 'IV', 'PhD', 'MD');

CREATE TABLE people (
    id bigserial PRIMARY KEY,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    prefix prefix,
    first_name text NOT NULL,
    middle_name text,
    last_name text NOT NULL,
    suffix suffix
);
