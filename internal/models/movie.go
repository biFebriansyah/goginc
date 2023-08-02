package models

import "time"

var schema = `
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
`

type Movie struct {
	Id_movie     string     `db:"id_movie" form:"id_movie" json:"id_movie"`
	Movie_name   string     `db:"movie_name" form:"movie_name" json:"movie_name"`
	Movie_banner string     `db:"movie_banner" form:"movie_banner" json:"movie_banner"`
	Release_date string     `db:"release_date" form:"release_date" json:"release_date"`
	CreatedAt    *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at" json:"updated_at"`
}
