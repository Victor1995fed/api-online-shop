CREATE TABLE product_tag (
    id bigserial not null  primary key,
    product_id int not null,
    tag_id int not null
);

create index idx_product_tag_product_id
    on product_tag (product_id);

create index idx_product_tag_tag_id
    on product_tag (tag_id);
