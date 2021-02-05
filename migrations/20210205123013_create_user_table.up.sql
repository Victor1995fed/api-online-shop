CREATE TABLE "user" (
    id bigserial not null  primary key,
    email varchar unique not null ,
    username varchar(255) unique not null,
    encrypted_password  varchar  not null
);

