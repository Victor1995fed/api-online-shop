CREATE TABLE "order_customer" (
    id bigserial not null  primary key,
    customer_id int default null,   
    order_id int default null
);

create index idx_order_customer_customer_id
    on order_customer (customer_id);

create index idx_order_customer_order_id
    on order_customer (order_id);