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
	User_id   string     `db:"user_id" json:"user_id,omitempty" form:"user_id" valid:"-"`
	Username  string     `db:"username" json:"username" form:"username" valid:"type(string)"`
	Password  string     `db:"password" json:"password,omitempty" valid:"stringlength(6|10)~Password minimal 6"`
	Role      string     `db:"role" json:"role,omitempty" valid:"-"`
	CreatedAt *time.Time `db:"created_at" json:"created_at" valid:"-"`
	UpdateAt  *time.Time `db:"updated_at" json:"updated_at" valid:"-"`
}

type Users []User
