CREATE TABLE IF NOT EXISTS cities (
    id int4 NOT NULL,
    id_province int4 NOT NULL,
    code varchar(18) NOT NULL,
    "name" varchar NOT NULL,
    postal_codes _int4 NULL,
    CONSTRAINT cities_pk PRIMARY KEY (id, id_province)
);