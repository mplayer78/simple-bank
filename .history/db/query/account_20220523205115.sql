-- name: GetAccount :one
SELECT * FROM Account
WHERE id = $1 LIMIT 1;

-- name: GetAccounts :many
SELECT * FROM Account
ORDER BY name;

-- name: CreateAccount :one
INSERT INTO Account (
  owner, balance, currency
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM Account
WHERE id = $1;