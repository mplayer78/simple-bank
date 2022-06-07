-- name: CreateAccount :one
INSERT INTO Accounts (
  owner, balance, currency
) VALUES (
  $1, $2, $3
)
RETURNING *;