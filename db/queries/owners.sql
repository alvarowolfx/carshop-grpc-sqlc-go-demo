-- name: CreateOwner :one
INSERT INTO owners (email, national_id)
VALUES ($1, $2)
RETURNING *;
-- name: GetOwnerByID :one
SELECT *
FROM owners
WHERE id = $1
LIMIT 1;
-- name: GetOwnerByEmail :one
SELECT *
FROM owners
WHERE email = $1
LIMIT 1;
-- name: GetOwnerByNationalID :one
SELECT *
FROM owners
WHERE national_id = $1
LIMIT 1;