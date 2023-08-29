CREATE TABLE tiketz.movie (
    movie_id uuid NULL DEFAULT gen_random_uuid(),
    movie_name varchar(255) NOT NULL,
    movie_banner varchar(255) NOT NULL,
    release_date date NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NULL,
    CONSTRAINT movie_pk PRIMARY KEY (movie_id)
);