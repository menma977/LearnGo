create table if not exists items
(
    id          serial primary key,
    uuid        uuid           not null default gen_random_uuid(),
    name        varchar(255)   not null,
    description text           null,
    price       numeric(10, 2) not null,
    quantity    bigint         not null,
    created_by  bigint         not null,
    updated_by  bigint         null,
    deleted_by  bigint         null,
    created_at  timestamptz    not null default now(),
    updated_at  timestamptz    not null default now(),
    deleted_at  timestamptz    null
)