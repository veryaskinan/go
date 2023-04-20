-- migrate:up
create table if not exists users (
    id uuid,
    phone text
);

-- migrate:down
drop table if exists users
