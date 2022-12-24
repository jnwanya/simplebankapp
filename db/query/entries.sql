-- name: CreateEntry :one
INSERT INTO entries (
    "account_id",
    "amount"
) VALUES (
     $1, $2
 ) RETURNING *;


-- name: GetEntry :one
select * from entries
where id = $1 limit 1;


-- name: ListEntries :many
select * from entries
order by id
limit $1
offset $2;
