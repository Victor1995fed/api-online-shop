CREATE TABLE product (
    id bigserial not null  primary key,
    title varchar not null,
    ddescription text,
    price decimal(19,4) default null,
    image_url varchar default null
)