package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomAssignment(t *testing.T) Assignment {
	ticket := CreateTicket(t)

	agent := CreateAgent(t)

	args := CreateAssignmentParams{
		TicketID: ticket.ID,
		AgentID:  agent.ID,
		Status:   "ASSIGNED",
	}

	assignment, err := testQuery.CreateAssignment(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, assignment)
	require.Equal(t, ticket.ID, assignment.TicketID)
	require.Equal(t, agent.ID, assignment.AgentID)
	require.Equal(t, args.Status, assignment.Status)
	require.NotZero(t, assignment.AssignedAt)

	return assignment
}

func TestCreateAssignment(t *testing.T) {
	createRandomAssignment(t)
}

func TestGetAssignment(t *testing.T) {
	newAssignment := createRandomAssignment(t)

	foundAssignment, err := testQuery.GetAssignment(context.Background(), newAssignment.ID)

	require.NoError(t, err)
	require.Equal(t, foundAssignment.ID, newAssignment.ID)
	require.Equal(t, newAssignment.TicketID, foundAssignment.TicketID)
	require.Equal(t, newAssignment.AgentID, foundAssignment.AgentID)
	require.WithinDuration(t, newAssignment.AssignedAt, foundAssignment.AssignedAt, time.Second)

}

func TestGetAssignmentForUpdate(t *testing.T) {
	newAssignment := createRandomAssignment(t)

	foundAssignment, err := testQuery.GetAssignment(context.Background(), newAssignment.ID)

	require.NoError(t, err)
	require.Equal(t, foundAssignment.ID, newAssignment.ID)
	require.Equal(t, newAssignment.TicketID, foundAssignment.TicketID)
	require.Equal(t, newAssignment.AgentID, foundAssignment.AgentID)
	require.WithinDuration(t, newAssignment.AssignedAt, foundAssignment.AssignedAt, time.Second)
}

func TestGetAssignments(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomAssignment(t)
	}

	params := GetAssignmentsParams{
		Limit:  5,
		Offset: 0,
	}

	assignments, err := testQuery.GetAssignments(context.Background(), params)

	require.NoError(t, err)
	require.NotEmpty(t, assignments)
	require.Len(t, assignments, 5)
}

func TestUpdateAssignment(t *testing.T) {
	newAssignment := createRandomAssignment(t)

	params := UpdateAssignmentParams{
		ID:     newAssignment.ID,
		Status: "RESOLVED",
	}

	updatedAssignment, err := testQuery.UpdateAssignment(context.Background(), params)

	require.NoError(t, err)
	require.NotEmpty(t, updatedAssignment)
	require.Equal(t, newAssignment.TicketID, updatedAssignment.TicketID)
	require.Equal(t, newAssignment.AgentID, updatedAssignment.AgentID)
	require.NotEqual(t, newAssignment.Status, updatedAssignment.Status)
	require.WithinDuration(t, newAssignment.AssignedAt, updatedAssignment.AssignedAt, time.Duration(time.Second))
}

func TestDeleteAssignment(t *testing.T) {
	newAssignment := createRandomAssignment(t)

	err := testQuery.DeleteAssignment(context.Background(), newAssignment.ID)

	require.NoError(t, err)

	foundAssignment, err := testQuery.GetAssignment(context.Background(), newAssignment.ID)
	require.Error(t, err)
	require.Empty(t, foundAssignment)
	require.EqualError(t, err, sql.ErrNoRows.Error())

}
