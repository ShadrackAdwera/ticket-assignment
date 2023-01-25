package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAgent(t *testing.T) {
	newAgent := CreateAgentParams{
		Name:   "Agent 007",
		Status: "ACTIVE",
	}
	res, err := testQuery.CreateAgent(context.Background(), newAgent)

	require.NoError(t, err)
	require.Equal(t, newAgent.Name, res.Name)
	require.Equal(t, newAgent.Status, res.Status)
	require.NotZero(t, res.CreatedAt)
}
