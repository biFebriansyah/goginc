package models

import "time"

var schemaUser = `
CREATE TABLE tiketz.user (
    user_id uuid NULL DEFAULT gen_random_uuid(),
    username varchar(255) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    role varchar(50) NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NULL,
    CONSTRAINT user_pk PRIMARY KEY (user_id)
);
`

type User struct {
	User_id   string     `db:"user_id" json:"user_id,omitempty" form:"user_id"`
	Username  string     `db:"username" json:"username" form:"username"`
	Password  string     `db:"password" json:"password,omitempty"`
	Role      string     `db:"role" json:"role,omitempty"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdateAt  *time.Time `db:"updated_at" json:"updated_at"`
}

type Users []User
