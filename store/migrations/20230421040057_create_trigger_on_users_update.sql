-- migrate:up
create function update_updated_at()
    returns trigger as $$
begin
    new.updated_at = now();
    return new;
end;
$$ language 'plpgsql';

drop trigger if exists
    users_update_at on users;

create trigger users_update_at
    after update on users
    for each row
execute procedure update_updated_at();


-- migrate:down
drop trigger if exists
    users_update_at on users;

drop function if exists update_updated_at;


