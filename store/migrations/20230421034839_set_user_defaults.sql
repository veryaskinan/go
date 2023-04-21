-- migrate:up
alter table if exists users
    alter column created_at set default now(),
    alter column updated_at set default now(),
    alter column active set default true
;

-- migrate:down
alter table if exists users
    alter column created_at drop default,
    alter column updated_at drop default,
    alter column active drop default
;

