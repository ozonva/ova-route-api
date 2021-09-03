-- +goose Up
-- +goose StatementBegin
CREATE table routes (
  id serial primary key,
  user_id serial,
  route_name text,
  length numeric
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table routes;
-- +goose StatementEnd
