-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id, to_account_id, amount
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: GetIncomingTransfers :many
SELECT * FROM transfers
WHERE to_account_id = $1
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetOutgoingTransfers :many
SELECT * FROM transfers
WHERE from_account_id = $1
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetBetweenTransfers :many
SELECT * FROM transfers
WHERE from_account_id = $1 AND to_account_id = $2
ORDER BY id
LIMIT $1
OFFSET $2;