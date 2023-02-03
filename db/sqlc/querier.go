// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"
)

type Querier interface {
	AssignTicketToAgent(ctx context.Context, arg AssignTicketToAgentParams) (Ticket, error)
	CreateAgent(ctx context.Context, arg CreateAgentParams) (Agent, error)
	CreateAssignment(ctx context.Context, arg CreateAssignmentParams) (Assignment, error)
	CreateTicket(ctx context.Context, arg CreateTicketParams) (Ticket, error)
	DeleteAgent(ctx context.Context, id int64) error
	DeleteAssignment(ctx context.Context, id int64) error
	DeleteTicket(ctx context.Context, id int64) error
	GetAgent(ctx context.Context, id int64) (Agent, error)
	GetAssignment(ctx context.Context, id int64) (GetAssignmentRow, error)
	GetAssignmentForUpdate(ctx context.Context, id int64) (Assignment, error)
	GetAssignments(ctx context.Context, arg GetAssignmentsParams) ([]GetAssignmentsRow, error)
	GetTicket(ctx context.Context, id int64) (Ticket, error)
	GetTicketForUpdate(ctx context.Context, id int64) (Ticket, error)
	GetTickets(ctx context.Context, arg GetTicketsParams) ([]Ticket, error)
	ListAgents(ctx context.Context, arg ListAgentsParams) ([]Agent, error)
	UpdateAgent(ctx context.Context, arg UpdateAgentParams) (Agent, error)
	UpdateAssignment(ctx context.Context, arg UpdateAssignmentParams) (Assignment, error)
	UpdateTicket(ctx context.Context, arg UpdateTicketParams) (Ticket, error)
}

var _ Querier = (*Queries)(nil)
