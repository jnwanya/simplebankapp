-- name: CreateTransfer :one
INSERT INTO transfers (
    "from_account_id",
    "to_account_id",
    "amount"
) VALUES (
     $1, $2, $3
 ) RETURNING *;


-- name: GetTransfer :one
select * from transfers
where id = $1 limit 1;


-- name: ListTransfers :many
select * from transfers
order by id
limit $1
offset $2;
