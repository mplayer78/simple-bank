-- name: CreateAccount :one
INSERT INTO Accounts (
  owner, balance, currency
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetAccount :one
SELECT * FROM Accounts
WHERE id = $1 LIMIT 1;

-- name: GetAllAccounts :many
SELECT * FROM Accounts
ORDER BY id
LIMIT $1
OFFSET $2;