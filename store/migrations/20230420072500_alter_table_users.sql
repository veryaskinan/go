-- migrate:up
alter table if exists users
    add column created_at timestamp,
    add column updated_at timestamp,
    add column active boolean
;

-- migrate:down
alter table if exists users
    drop column created_at,
    drop column updated_at,
    drop column active
;

