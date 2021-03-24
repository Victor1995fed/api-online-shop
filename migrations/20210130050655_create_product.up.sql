CREATE TABLE product (
    id bigserial not null  primary key,
    title varchar not null,
    description text,
    price decimal(19,4) default null
)