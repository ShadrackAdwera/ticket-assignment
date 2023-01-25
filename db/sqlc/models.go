// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
	"time"
)

type Agent struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	// Agent Status can be active or inactive
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type Assignment struct {
	ID       int64 `json:"id"`
	TicketID int64 `json:"ticket_id"`
	AgentID  int64 `json:"agent_id"`
	// Ticket Status can be PENDING or RESOLVED
	Status     string    `json:"status"`
	AssignedAt time.Time `json:"assigned_at"`
}

type Ticket struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	// Ticket Status can be PENDING or RESOLVED
	Status     string        `json:"status"`
	AssignedTo sql.NullInt64 `json:"assigned_to"`
	CreatedAt  time.Time     `json:"created_at"`
}
