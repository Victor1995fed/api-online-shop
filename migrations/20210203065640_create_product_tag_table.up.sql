CREATE TABLE product_tag (
    id bigserial not null  primary key,
    id_product int not null,
    id_tag int not null
)