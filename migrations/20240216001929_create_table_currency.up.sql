CREATE TABLE currency (
    id bigserial not null primary key,
    title character varying(60) not null,
    code character varying(3),
    value numeric(18, 2),
    date timestamptz,
    CONSTRAINT unique_code_date UNIQUE (code, date)
);