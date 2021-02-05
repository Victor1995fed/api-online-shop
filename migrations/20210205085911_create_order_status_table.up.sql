CREATE TABLE order_status (
    id bigserial not null  primary key,
    short_name varchar(255) unique,   
    title varchar(255)
);