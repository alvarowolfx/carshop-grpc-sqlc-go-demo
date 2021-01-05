-- name: RegisterWorkOrderServiceExecution :one
INSERT INTO work_order_service_executions (type, work_order_id)
VALUES ($1, $2)
RETURNING *;
-- name: EndWorkOrderServiceExecution :exec
UPDATE work_order_service_executions
SET finished_at = NOW()
WHERE work_order_id = $1
  and finished_at is null;
-- name: GetRunningServices :many
SELECT *
FROM work_order_service_executions
WHERE work_order_id = $1
  and finished_at is null;