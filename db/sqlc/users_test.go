package db

import (
	"context"
	"jnwanya/simplebank/db/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {

	hashedPassword, err := util.HashedPassword("p@ssw0rd")
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)

	require.Equal(t, arg.HashedPassword, user.HashedPassword)

	require.Equal(t, arg.Email, user.Email)

	require.Equal(t, arg.FullName, user.FullName)

	require.True(t, user.PasswordChangedAt.IsZero())

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)

	user2, err := testQueries.GetUser(context.Background(), user.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user.Username, user2.Username)
	require.Equal(t, user.FullName, user2.FullName)
	require.Equal(t, user.HashedPassword, user2.HashedPassword)
	require.Equal(t, user.Email, user2.Email)
	require.WithinDuration(t, user.CreatedAt, user2.CreatedAt, time.Second)

}