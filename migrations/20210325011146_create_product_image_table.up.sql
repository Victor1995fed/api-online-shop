CREATE TABLE product_image (
    id bigserial not null  primary key,
    product_id int not null,
    image_id int not null
);

create index idx_product_image_product_id
    on product_image (product_id);

create index idx_product_image_image_id
    on product_image (image_id);