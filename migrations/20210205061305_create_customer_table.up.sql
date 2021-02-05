CREATE TABLE customer (
    id bigserial not null  primary key,
    date_create timestamp default now(),
    date_update timestamp default null,
    first_name varchar(255) not null, 
    last_name varchar(255),
    patronymic varchar(255) default null,
    email varchar(100) unique default null,
    phone    bigint unique default null,
    encrypted_password varchar  not null,
    address text
);

create index idx_customer_first_name
    on customer (first_name);

create index idx_customer_email
    on customer (email);

create index idx_customer_phone
    on customer (phone);