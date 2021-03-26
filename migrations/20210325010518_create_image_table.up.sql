create table image
(
    id   bigserial not null  primary key,
    hash varchar(36) null,
    constraint hash
        unique (hash)
);

create index idx_image_hash
    on image (hash);