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
		Username:       "testUser",
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

func TestGetUserById(t *testing.T) {

	user, err := testQueries.GetUserById(context.Background(), createdUserId)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.ID, createdUserId)
	require.Equal(t, user.Username, "testUser")
	require.Equal(t, user.HashedPassword, "123445544123")

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	require.Equal(t, user.ID, createdUserId)

}
func TestGetUserByUsername(t *testing.T) {

	user, err := testQueries.GetUserByUsername(context.Background(), "testUser")
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.Username, "testUser")
	require.Equal(t, user.HashedPassword, "123445544123")

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	require.Equal(t, user.ID, createdUserId)

}

func TestUpdateUser(t *testing.T) {
	user, err := testQueries.GetUserByUsername(context.Background(), "testUser")
	require.NoError(t, err)
	require.NotZero(t, user.ID)

	arg := UpdateUserParams{
		Username: sql.NullString{
			String: "testUser2",
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

	gotUser, err := testQueries.GetUserByUsername(context.Background(), "testUser2")
	require.NoError(t, err)
	require.NotEmpty(t, gotUser)

	require.Equal(t, gotUser.Username, "testUser2")

}

func TestDeleteUser(t *testing.T) {
	err := testQueries.DeleteUser(context.Background(), createdUserId)
	require.NoError(t, err)

	_, err = testQueries.GetUserByUsername(context.Background(), "testUser2")
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())

}
