CREATE TYPE roles_enum AS ENUM (
    'guest',
    'candidate',
    'rescuer',
    'operator',
    'admin'
);

CREATE TABLE IF NOT EXISTS users (
    id_user integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    login varchar(255) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    role roles_enum NOT NULL DEFAULT 'guest'
);
