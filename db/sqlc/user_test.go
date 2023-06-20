package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

var createdUserId int64

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		Username:       "mehmettest",
		HashedPassword: "123445544123",
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	createdUserId = user.ID
}

func TestGetUser(t *testing.T) {

	user, err := testQueries.GetUser(context.Background(), "mehmettest")
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.Username, "mehmettest")
	require.Equal(t, user.HashedPassword, "123445544123")

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	require.Equal(t, user.ID, createdUserId)

}

func TestUpdateUser(t *testing.T) {
	user, err := testQueries.GetUser(context.Background(), "mehmettest")
	require.NoError(t, err)
	require.NotZero(t, user.ID)

	arg := UpdateUserParams{
		Username: sql.NullString{
			String: "mehmettest2",
			Valid:  true,
		},
		HashedPassword: sql.NullString{
			String: "123456",
			Valid:  true,
		},
		ID: user.ID,
	}

	updatedUser, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedUser)

	gotUser, err := testQueries.GetUser(context.Background(), "mehmettest2")
	require.NoError(t, err)
	require.NotEmpty(t, gotUser)

	require.Equal(t, gotUser.Username, "mehmettest2")

}

func TestDeleteUser(t *testing.T) {
	err := testQueries.DeleteUser(context.Background(), createdUserId)
	require.NoError(t, err)

	_, err = testQueries.GetUser(context.Background(), "mehmettest2")
	require.Error(t, err)

}
