
-- name: CreateUsers :execresult
insert into
    users(username, password)
values
    ( $1, $2);

-- name: GetUserByName :one
select
    *
from
    users
where
    username =  $1;

-- name: GetAllUsers :many
select * from users;