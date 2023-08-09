CREATE TABLE tiketz.user (
    user_id uuid NULL DEFAULT gen_random_uuid(),
    username varchar(255) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    role varchar(50) NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NULL,
    CONSTRAINT user_pk PRIMARY KEY (user_id)
);