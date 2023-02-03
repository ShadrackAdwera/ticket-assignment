package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewTicketAssignment(t *testing.T) {
	txStore := NewTxStore(testDb)

	ticket1 := CreateTicket(t)
	//ticket2:= CreateTicket(t)

	agent1 := CreateAgent(t)
	//agent2 := CreateAgent(t)

	n := 5
	status := "ASSIGNED"
	errChan := make(chan error)
	resultChan := make(chan NewTicketAssignmentResult)

	for i := 0; i < n; i++ {
		go func() {
			res, err := txStore.NewTicketAssignment(context.Background(), CreateAssignmentParams{
				TicketID: ticket1.ID,
				AgentID:  agent1.ID,
				Status:   status,
			})
			errChan <- err
			resultChan <- res
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errChan
		require.NoError(t, err)

		result := <-resultChan

		require.NotEmpty(t, result)
		require.Equal(t, result.Agent.ID, agent1.ID)
		require.Equal(t, result.Ticket.ID, ticket1.ID)
		require.NotZero(t, result.Assignment.AssignedAt)
		require.Equal(t, result.Assignment.Status, status)

		dbCol, err := txStore.GetAssignment(context.Background(), result.Assignment.ID)

		require.NoError(t, err)
		require.Equal(t, dbCol.ID, result.Assignment.ID)
		require.Equal(t, dbCol.TicketID, result.Assignment.TicketID)
		require.Equal(t, dbCol.AgentID, result.Assignment.AgentID)
		require.WithinDuration(t, dbCol.AssignedAt, result.Assignment.AssignedAt, time.Second)

	}

}
