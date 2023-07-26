// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: posts.sql

package db

import (
	"context"
	"database/sql"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (
    user_id,
    title,
    body,
    status
) VALUES (
    $1, $2, $3, $4
         ) RETURNING id, user_id, title, body, status, created_at, updated_at
`

type CreatePostParams struct {
	UserID int64  `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Status string `json:"status"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.UserID,
		arg.Title,
		arg.Body,
		arg.Status,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Body,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deletePost = `-- name: DeletePost :exec
DELETE FROM posts
WHERE
        id = $1
`

func (q *Queries) DeletePost(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePost, id)
	return err
}

const getPostById = `-- name: GetPostById :one
SELECT id, user_id, title, body, status, created_at, updated_at FROM posts
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPostById(ctx context.Context, id int64) (Post, error) {
	row := q.db.QueryRowContext(ctx, getPostById, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Body,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPostsByUserId = `-- name: GetPostsByUserId :many
SELECT id, user_id, title, body, status, created_at, updated_at FROM posts
WHERE user_id = $1
`

func (q *Queries) GetPostsByUserId(ctx context.Context, userID int64) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getPostsByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Body,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPosts = `-- name: ListPosts :many
SELECT id, user_id, title, body, status, created_at, updated_at FROM posts
WHERE status = $1 LIMIT $2 OFFSET $3
`

type ListPostsParams struct {
	Status string `json:"status"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

func (q *Queries) ListPosts(ctx context.Context, arg ListPostsParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, listPosts, arg.Status, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Body,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePost = `-- name: UpdatePost :one
UPDATE posts
SET
    user_id = COALESCE($1, user_id),
    title = COALESCE($2, title),
    body = COALESCE($3, body),
    status = COALESCE($4, status),
    updated_at = now()
WHERE
        id = $5
RETURNING id, user_id, title, body, status, created_at, updated_at
`

type UpdatePostParams struct {
	UserID sql.NullInt64  `json:"user_id"`
	Title  sql.NullString `json:"title"`
	Body   sql.NullString `json:"body"`
	Status sql.NullString `json:"status"`
	ID     int64          `json:"id"`
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, updatePost,
		arg.UserID,
		arg.Title,
		arg.Body,
		arg.Status,
		arg.ID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Body,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
