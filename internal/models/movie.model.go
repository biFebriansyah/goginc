package models

import "time"

var schemaMovie = `
CREATE TABLE tiketz.movie (
    movie_id uuid NULL DEFAULT gen_random_uuid(),
    movie_name varchar(255) NOT NULL,
    movie_banner varchar(255) NOT NULL,
    release_date date NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NULL,
    CONSTRAINT movie_pk PRIMARY KEY (movie_id)
);
`

type Movie struct {
	Movie_id     string     `db:"movie_id" json:"movie_id,omitempty" form:"movie_id"`
	Movie_name   string     `db:"movie_name" json:"movie_name" form:"movie_name"`
	Movie_banner string     `db:"movie_banner" json:"movie_banner,omitempty"`
	Release_date *string    `db:"release_date" json:"release_date" form:"release_date"`
	CreatedAt    *time.Time `db:"created_at" json:"created_at"`
	UpdateAt     *time.Time `db:"updated_at" json:"updated_at"`
}

type Meta struct {
	Page  int
	Limit int
	Name  string
}

type Movies []Movie
