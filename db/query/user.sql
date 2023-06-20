-- name: CreateUser :one
INSERT INTO users (
    username,
    hashed_password
) VALUES (
             $1, $2
         ) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET
    hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
    username = COALESCE(sqlc.narg(username), username)
WHERE
        id = sqlc.arg(id)
    RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE
    id = $1;
