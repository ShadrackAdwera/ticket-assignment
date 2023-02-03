package db

import (
	"context"
	"fmt"

	"testing"
	"time"

	"github.com/ShadrackAdwera/ticket-assignment/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	username := utils.RandomString(10)

	newUser := CreateUserParams{
		Username: username,
		Email:    fmt.Sprintf("%s@mail.com", username),
		Password: utils.RandomString(10),
	}

	user, err := testQuery.CreateUser(context.Background(), newUser)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.NotZero(t, user.ID)
	require.Equal(t, newUser.Username, user.Username)
	require.Equal(t, newUser.Email, user.Email)
	require.NotZero(t, user.CreatedAt)
	require.NotEmpty(t, user.Password)
	require.True(t, user.PasswordChangedAt.IsZero())

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	newUser := createRandomUser(t)

	user, err := testQuery.GetUser(context.Background(), newUser.ID)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, newUser.Username, user.Username)
	require.Equal(t, newUser.Email, user.Email)
	require.WithinDuration(t, user.CreatedAt, newUser.CreatedAt, time.Second)
	require.NotEmpty(t, user.Password)
	require.WithinDuration(t, user.PasswordChangedAt, newUser.PasswordChangedAt, time.Second)
}

func TestGetUsers(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomUser(t)
	}

	users, err := testQuery.ListUsers(context.Background(), ListUsersParams{
		Limit:  5,
		Offset: 1,
	})

	require.NoError(t, err)
	require.NotEmpty(t, users)
}
