-- +migrate Up
CREATE TABLE owners (
  id SERIAL PRIMARY KEY,
  email TEXT NOT NULL UNIQUE,
  national_id TEXT NOT NULL UNIQUE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- +migrate StatementBegin
CREATE TRIGGER set_owners_updated_at BEFORE
UPDATE ON owners FOR EACH ROW EXECUTE PROCEDURE set_current_timestamp_updated_at();
-- +migrate StatementEnd
-- +migrate Down
DROP TABLE owners;