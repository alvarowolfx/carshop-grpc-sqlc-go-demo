-- +migrate Up
CREATE TYPE car_size as ENUM ('small', 'medium', 'large');
CREATE TABLE cars (
  id SERIAL PRIMARY KEY,
  license_plate TEXT NOT NULL UNIQUE,
  size car_size NOT NULL,
  num_wheels SMALLINT NOT NULL,
  color TEXT NOT NULL,
  owner_id BIGINT REFERENCES owners (id) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- +migrate StatementBegin
CREATE TRIGGER set_cars_updated_at BEFORE
UPDATE ON cars FOR EACH ROW EXECUTE PROCEDURE set_current_timestamp_updated_at();
-- +migrate StatementEnd
-- +migrate Down
DROP TABLE cars;
DROP TYPE car_size;