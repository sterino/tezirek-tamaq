-- +goose Up
-- +goose StatementBegin
CREATE TABLE restaurants (
                          id         UUID PRIMARY KEY,
                          name       TEXT NOT NULL,
                          address    TEXT,
                          phone      TEXT,
                          order_ids  TEXT[] DEFAULT '{}', -- массив строк
                          created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
                          updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE restaurants;
-- +goose StatementEnd
