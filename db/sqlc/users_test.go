package db

import (
	"context"
	"testing"

	"github.com/karan0033/bank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := utils.HashPassword(utils.RandomString(8))
	require.NoError(t, err)
	arg := CreateUsersParams{
		Username:       utils.RandomOwner(),
		FullName:       utils.RandomOwner(),
		HashedPassword: hashedPassword,
		Email:          utils.RandomEmail(),
	}

	user, err := testQuery.CreateUsers(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.Email, user.Email)

	require.NotEmpty(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)

	user2, err := testQuery.GetUser(context.Background(), user1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user2.Username, user1.Username)
	require.Equal(t, user2.FullName, user1.FullName)
	require.Equal(t, user2.Email, user1.Email)
	require.Equal(t, user2.HashedPassword, user1.HashedPassword)
	require.Equal(t, user2.CreatedAt, user1.CreatedAt)
}
