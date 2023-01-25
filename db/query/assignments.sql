-- name: GetAssignment :one
SELECT * FROM assignments
WHERE id = $1 LIMIT 1;

-- name: GetAssignments :many
SELECT * FROM assignments
ORDER BY id;

-- name: GetAssignmentForUpdate :one
SELECT * FROM assignments
WHERE id = $1 LIMIT 1 
FOR NO KEY UPDATE;

-- name: CreateAssignment :one
INSERT INTO assignments (
ticket_id, agent_id, status
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateAssignment :one
UPDATE assignments SET status = $2
WHERE id = $1 RETURNING *;

-- name: DeleteAssignment :exec
DELETE FROM assignments
WHERE id = $1;