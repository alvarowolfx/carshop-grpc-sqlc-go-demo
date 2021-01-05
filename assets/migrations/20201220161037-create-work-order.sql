-- +migrate Up
CREATE TYPE work_order_status as ENUM (
  'CREATED',
  'DIAGNOSTICS',
  'CHANGING_PARTS',
  'CHANGING_TIRES',
  'WASHING',
  'IDLE',
  'FINISHED',
  'DONE'
);
CREATE TABLE work_orders (
  id SERIAL PRIMARY KEY,
  change_tires boolean NOT NULL,
  change_parts boolean NOT NULL,
  current_status work_order_status NOT NULL,
  previous_status work_order_status,
  car_id BIGINT REFERENCES cars (id) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- +migrate StatementBegin
CREATE TRIGGER set_work_orders_updated_at BEFORE
UPDATE ON work_orders FOR EACH ROW EXECUTE PROCEDURE set_current_timestamp_updated_at();
-- +migrate StatementEnd
-- +migrate Down
DROP TABLE work_orders;
DROP TYPE work_order_status;