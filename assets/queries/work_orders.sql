-- name: CreateWorkOrder :one
INSERT INTO work_orders (
    change_tires,
    change_parts,
    car_id,
    previous_status,
    current_status
  )
VALUES ($1, $2, $3, 'CREATED', 'CREATED')
RETURNING *;
-- name: GetRunningWorkOrders :many
SELECT *
FROM work_orders wo
  inner join cars c on wo.car_id = c.id
WHERE wo.current_status not in ('DONE');
-- name: UpdateWorkOrderServiceStatus :exec
UPDATE work_orders
SET previous_status = current_status,
  current_status = $2
WHERE id = $1;
-- name: EndWorkOrder :exec
UPDATE work_orders
SET previous_status = current_status,
  current_status = 'DONE'
WHERE id = $1;