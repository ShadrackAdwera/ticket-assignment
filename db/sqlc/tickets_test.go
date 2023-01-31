package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ShadrackAdwera/ticket-assignment/utils"
	"github.com/stretchr/testify/require"
)

func CreateTicket(t *testing.T) Ticket {
	newTicketArgs := CreateTicketParams{
		Title:       utils.GenerateTicketTitle(),
		Description: utils.GenerateTicketDescription(),
	}

	ticket, err := testQuery.CreateTicket(context.Background(), newTicketArgs)

	require.NoError(t, err)
	require.NotZero(t, ticket.CreatedAt)
	require.Equal(t, newTicketArgs.Title, ticket.Title)
	require.Equal(t, newTicketArgs.Description, ticket.Description)

	return ticket
}

func TestCreateTicket(t *testing.T) {
	CreateTicket(t)
}

func TestGetTicket(t *testing.T) {
	newTicket := CreateTicket(t)

	ticket, err := testQuery.GetTicket(context.Background(), newTicket.ID)

	require.NoError(t, err)
	require.NotEmpty(t, ticket)
	require.Equal(t, newTicket.Title, ticket.Title)
	require.Equal(t, newTicket.Description, ticket.Description)
	require.WithinDuration(t, newTicket.CreatedAt, ticket.CreatedAt, time.Duration(time.Second))

	ticc, err := testQuery.GetTicket(context.Background(), 91928273737211)

	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ticc)
}

func TestGetTickets(t *testing.T) {
	for i := 0; i < 5; i++ {
		CreateTicket(t)
	}

	params := GetTicketsParams{
		Limit:  5,
		Offset: 1,
	}

	tickets, err := testQuery.GetTickets(context.Background(), params)

	require.NoError(t, err)
	require.NotEmpty(t, tickets)
	require.Len(t, tickets, 5)
}

func TestUpdateTicket(t *testing.T) {
	ticket := CreateTicket(t)

	updateParams := UpdateTicketParams{
		ID: ticket.ID,
	}
	updatedTicket, err := testQuery.UpdateTicket(context.Background(), updateParams)

	require.NoError(t, err)
	require.Equal(t, ticket.Title, updatedTicket.Title)
	require.Equal(t, ticket.Description, updatedTicket.Description)
	require.WithinDuration(t, ticket.CreatedAt, updatedTicket.CreatedAt, time.Second)
	require.NotEqual(t, updatedTicket.Status, ticket.Status)
}

func TestAssignTicketToAgent(t *testing.T) {
	ticket := CreateTicket(t)
	agent := CreateAgent(t)

	params := AssignTicketToAgentParams{
		ID: ticket.ID,
		AssignedTo: sql.NullInt64{
			Int64: agent.ID,
			Valid: true,
		},
		Status: "IN PROGRESS",
	}
	assignedTicket, err := testQuery.AssignTicketToAgent(context.Background(), params)

	require.NoError(t, err)
	require.Equal(t, ticket.Title, assignedTicket.Title)
	require.Equal(t, ticket.Description, assignedTicket.Description)
	require.WithinDuration(t, ticket.CreatedAt, assignedTicket.CreatedAt, time.Second)
	require.NotEqual(t, assignedTicket.Status, ticket.Status)
	require.Equal(t, sql.NullInt64{
		Int64: agent.ID,
		Valid: true,
	}, assignedTicket.AssignedTo)
}

func TestDeleteTicket(t *testing.T) {
	ticket := CreateTicket(t)

	err := testQuery.DeleteTicket(context.Background(), ticket.ID)
	require.NoError(t, err)

	foundTicket, err := testQuery.GetTicket(context.Background(), ticket.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, foundTicket)
}
