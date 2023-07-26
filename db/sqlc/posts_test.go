package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatePost(t *testing.T) {
	// create new user
	arg := CreateUserParams{
		Username:       "testPostUser",
		HashedPassword: "123445544123",
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	createdUserId := user.ID

	// create post
	arg2 := CreatePostParams{
		UserID: createdUserId,
		Title:  "test post title",
		Body:   "test post body",
		Status: "active",
	}

	post, err := testQueries.CreatePost(context.Background(), arg2)
	require.NoError(t, err)
	require.NotEmpty(t, post)
}
