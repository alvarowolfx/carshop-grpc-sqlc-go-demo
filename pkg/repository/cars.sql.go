// Code generated by sqlc. DO NOT EDIT.
// source: cars.sql

package repository

import (
	"context"
)

const createCar = `-- name: CreateCar :one
INSERT INTO cars (license_plate, size, num_wheels, color, owner_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, license_plate, size, num_wheels, color, owner_id, created_at, updated_at
`

type CreateCarParams struct {
	LicensePlate string  `json:"license_plate"`
	Size         CarSize `json:"size"`
	NumWheels    int16   `json:"num_wheels"`
	Color        string  `json:"color"`
	OwnerID      int64   `json:"owner_id"`
}

func (q *Queries) CreateCar(ctx context.Context, arg CreateCarParams) (Car, error) {
	row := q.queryRow(ctx, q.createCarStmt, createCar,
		arg.LicensePlate,
		arg.Size,
		arg.NumWheels,
		arg.Color,
		arg.OwnerID,
	)
	var i Car
	err := row.Scan(
		&i.ID,
		&i.LicensePlate,
		&i.Size,
		&i.NumWheels,
		&i.Color,
		&i.OwnerID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCarByLicensePlate = `-- name: GetCarByLicensePlate :one
SELECT id, license_plate, size, num_wheels, color, owner_id, created_at, updated_at
FROM cars
WHERE license_plate = $1
`

func (q *Queries) GetCarByLicensePlate(ctx context.Context, licensePlate string) (Car, error) {
	row := q.queryRow(ctx, q.getCarByLicensePlateStmt, getCarByLicensePlate, licensePlate)
	var i Car
	err := row.Scan(
		&i.ID,
		&i.LicensePlate,
		&i.Size,
		&i.NumWheels,
		&i.Color,
		&i.OwnerID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCarsByOwner = `-- name: GetCarsByOwner :many
SELECT id, license_plate, size, num_wheels, color, owner_id, created_at, updated_at
FROM cars
WHERE owner_id = $1
`

func (q *Queries) GetCarsByOwner(ctx context.Context, ownerID int64) ([]Car, error) {
	rows, err := q.query(ctx, q.getCarsByOwnerStmt, getCarsByOwner, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Car
	for rows.Next() {
		var i Car
		if err := rows.Scan(
			&i.ID,
			&i.LicensePlate,
			&i.Size,
			&i.NumWheels,
			&i.Color,
			&i.OwnerID,
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
