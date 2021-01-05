-- name: CreateCar :one
INSERT INTO cars (license_plate, size, num_wheels, color, owner_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: GetCarsByOwner :many
SELECT *
FROM cars
WHERE owner_id = $1;
-- name: GetCarByLicensePlate :one
SELECT *
FROM cars
WHERE license_plate = $1;