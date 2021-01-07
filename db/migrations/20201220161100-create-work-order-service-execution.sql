-- +migrate Up
CREATE TYPE service_type as ENUM (
  'DIAGNOSTIC',
  'CHANGE_PARTS',
  'CHANGE_TIRES',
  'WASH'
);
CREATE TABLE work_order_service_executions (
  id SERIAL PRIMARY KEY,
  type service_type NOT NULL,
  work_order_id BIGINT REFERENCES work_orders (id) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  finished_at TIMESTAMPTZ
);
-- +migrate Down
DROP TABLE work_order_service_executions;
DROP TYPE service_type;