-- name: CreateAccount :one
INSERT INTO accounts (
    owner,
    balance,
    currency
) VALUES (
    $1, $2, $3
 ) RETURNING *;


-- name: GetAccount :one
select * from accounts
where id = $1 limit 1;

-- name: GetAccountForUpdate :one
select * from accounts
where id = $1 limit 1 
FOR NO KEY UPDATE;

-- name: AddAccountBalance :one
update accounts
set balance = balance + sqlc.arg(amount)
where id = sqlc.arg(id)
RETURNING *;

-- name: ListAccounts :many
select * from accounts
order by id
limit $1
offset $2;


-- name: UpdateAccount :one
update accounts
set balance = $2
where id = $1
RETURNING *;

-- name: DeleteAccount :exec
delete from accounts
where id = $1;
