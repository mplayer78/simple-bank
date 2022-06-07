-- name: CreateEntry :one
INSERT INTO Entries (
  account_id, amount
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetEntry :one
SELECT * FROM Entries
WHERE id = $1 LIMIT 1;

-- name: GetAllAccountEntries :many
SELECT * FROM Entries
WHERE account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;