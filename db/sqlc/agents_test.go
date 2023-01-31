package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ShadrackAdwera/ticket-assignment/utils"
	"github.com/stretchr/testify/require"
)

func CreateAgent(t *testing.T) Agent {
	agentArgs := CreateAgentParams{
		Name:   utils.GenerateRandomString(),
		Status: utils.GetAgentStatus(),
	}
	agent, err := testQuery.CreateAgent(context.Background(), agentArgs)

	require.NoError(t, err)
	require.NotEmpty(t, agent)
	require.Equal(t, agentArgs.Name, agent.Name)
	require.Equal(t, agentArgs.Status, agent.Status)
	require.NotZero(t, agent.CreatedAt)

	return agent
}

func TestCreateAgent(t *testing.T) {
	newAgent := CreateAgentParams{
		Name:   utils.GenerateRandomString(),
		Status: utils.GetAgentStatus(),
	}
	res, err := testQuery.CreateAgent(context.Background(), newAgent)

	require.NoError(t, err)
	require.Equal(t, newAgent.Name, res.Name)
	require.Equal(t, newAgent.Status, res.Status)
	require.NotZero(t, res.CreatedAt)
}

func TestListAgents(t *testing.T) {
	for i := 0; i < 5; i++ {
		CreateAgent(t)
	}

	listAgentParams := ListAgentsParams{
		Limit:  5,
		Offset: 1,
	}

	agents, err := testQuery.ListAgents(context.Background(), listAgentParams)

	require.NoError(t, err)
	require.NotEmpty(t, agents)
	require.Len(t, agents, 5)
}

func TestGetAgent(t *testing.T) {
	createdAgent := CreateAgent(t)
	foundAgent, err := testQuery.GetAgent(context.Background(), createdAgent.ID)

	require.NoError(t, err)
	require.NotEmpty(t, foundAgent)
	require.Equal(t, createdAgent.Name, foundAgent.Name)
	require.Equal(t, createdAgent.Status, foundAgent.Status)
	require.NotZero(t, createdAgent.CreatedAt)

	foundAgent, err = testQuery.GetAgent(context.Background(), 13242225252525252)
	require.ErrorIs(t, err, sql.ErrNoRows)
	require.Empty(t, foundAgent)
}

func TestUpdateAgent(t *testing.T) {
	createdAgent := CreateAgent(t)

	updatedAgent := UpdateAgentParams{
		ID:     createdAgent.ID,
		Status: utils.GetAgentStatus(),
	}

	agent, err := testQuery.UpdateAgent(context.Background(), updatedAgent)

	require.NoError(t, err)
	require.Equal(t, updatedAgent.Status, agent.Status)
	require.Equal(t, createdAgent.Name, agent.Name)
	require.WithinDuration(t, agent.CreatedAt, createdAgent.CreatedAt, time.Duration(time.Second*1))
}

func TestDeleteAgent(t *testing.T) {
	createdAgent := CreateAgent(t)

	err := testQuery.DeleteAgent(context.Background(), createdAgent.ID)
	require.NoError(t, err)

	foundAgent, err := testQuery.GetAgent(context.Background(), createdAgent.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, foundAgent)
}
