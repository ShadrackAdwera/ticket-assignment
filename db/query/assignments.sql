-- name: GetAssignment :one
SELECT * FROM assignments
JOIN tickets ON assignments.ticket_id = tickets.id
JOIN agents ON assignments.agent_id = agents.id
WHERE assignments.id = $1 
LIMIT 1;

-- name: GetAssignments :many
SELECT * FROM assignments
JOIN tickets ON assignments.ticket_id = tickets.id
JOIN agents ON assignments.agent_id = agents.id
ORDER BY assignments.id
LIMIT $1
OFFSET $2;

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