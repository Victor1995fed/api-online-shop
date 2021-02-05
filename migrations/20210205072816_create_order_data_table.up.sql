CREATE TABLE order_data (
    id bigserial not null  primary key,
    address text,   
    phone bigint default null,
    order_id int,
    email varchar(100) default null,
    first_name varchar(255),
    last_name varchar(255),
    patronymic varchar(255) default null,
    comment text
);

create index idx_order_data_phone
    on order_data (phone);

create index idx_order_data_email
    on order_data (email);