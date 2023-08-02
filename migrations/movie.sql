CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE tiketz.movie (
    id_movie uuid NULL DEFAULT uuid_generate_v4(),
    movie_name varchar(255) NOT NULL,
    slug_movie varchar(255) NOT NULL UNIQUE,
    movie_banner varchar(255) NOT NULL,
    release_date date NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NULL,
    CONSTRAINT movie_pk PRIMARY KEY (id_movie)
);