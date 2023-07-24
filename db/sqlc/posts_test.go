package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatePost(t *testing.T) {
	arg := CreatePostParams{
		UserID: 1,
		Title:  "test post title",
		Body:   "test post body",
		Status: "active",
	}

	post, err := testQueries.CreatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post)
}
