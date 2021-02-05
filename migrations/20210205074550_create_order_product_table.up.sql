CREATE TABLE order_product (
    id bigserial not null  primary key,
    product_id int default null,   
    order_id int default null
);

create index idx_order_product_product_id
    on order_product (product_id);

create index idx_order_product_order_id
    on order_product (order_id);