-- name: CreateAccount :one
INSERT INTO Account (
  owner, balance, currency
) VALUES (
  $1, $2, $3
)
RETURNING *;