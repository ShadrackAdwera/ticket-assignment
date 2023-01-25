-- name: GetAgent :one
SELECT * FROM agents
WHERE id = $1 LIMIT 1;

-- name: ListAgents :many
SELECT * FROM agents
ORDER BY name
LIMIT $1
OFFSET $2;

-- name: CreateAgent :one
INSERT INTO agents (
  name, status
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateAgent :one
UPDATE agents SET status = $2
WHERE id = $1 RETURNING *;

-- name: DeleteAgent :exec
DELETE FROM agents
WHERE id = $1;