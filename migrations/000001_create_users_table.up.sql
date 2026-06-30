create table if not exists users
(
    id         serial primary key,
    uuid       uuid         not null default gen_random_uuid(),
    name       varchar(255) not null,
    username   varchar(255) not null,
    email      varchar(255) not null,
    password   varchar(255) not null,
    created_at timestamptz  not null default now(),
    updated_at timestamptz  not null default now(),
    deleted_at timestamp    null
);

create unique index users_uuid_index on users (uuid);