-- name: CreatePost :one
INSERT INTO posts (
    user_id,
    title,
    body,
    status
) VALUES (
    $1, $2, $3, $4
         ) RETURNING *;

-- name: GetPostById :one
SELECT * FROM posts
WHERE id = $1 LIMIT 1;

-- name: ListPosts :many
SELECT * FROM posts
WHERE status = $1 LIMIT $2 OFFSET $3;

-- name: GetPostsByUserId :many
SELECT * FROM posts
WHERE user_id = $1;

-- name: UpdatePost :one
UPDATE posts
SET
    user_id = COALESCE(sqlc.narg(user_id), user_id),
    title = COALESCE(sqlc.narg(title), title),
    body = COALESCE(sqlc.narg(body), body),
    status = COALESCE(sqlc.narg(status), status),
    updated_at = now()
WHERE
        id = sqlc.arg(id)
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts
WHERE
        id = $1;
