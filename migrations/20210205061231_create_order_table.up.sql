CREATE TABLE "order" (
    id bigserial not null  primary key,
    date_create timestamp default now(),
    date_update timestamp default null,
    status_id int default null
);


create index idx_order_status_id
    on "order" (status_id);