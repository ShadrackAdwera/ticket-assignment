-- name: GetTicket :one
SELECT * FROM tickets
WHERE id = $1 LIMIT 1;

-- name: GetTickets :many
SELECT * FROM tickets
LIMIT $1 OFFSET $2;

-- name: GetTicketForUpdate :one
SELECT * FROM tickets
WHERE id = $1 LIMIT 1 
FOR NO KEY UPDATE;

-- name: CreateTicket :one
INSERT INTO tickets (
  title, description, createdby_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateTicket :one
UPDATE tickets SET status = $2
WHERE id = $1 RETURNING *;

-- name: AssignTicketToAgent :one
UPDATE tickets SET assigned_to = $2, 
status = $3
WHERE id = $1 RETURNING *;

-- name: DeleteTicket :exec
DELETE FROM tickets
WHERE id = $1;