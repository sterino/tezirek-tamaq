-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
                          id             UUID PRIMARY KEY,
                          user_id        UUID NOT NULL,
                          restaurant_id  UUID NOT NULL,
                          items          JSONB NOT NULL,
                          total_price    NUMERIC(10, 2) NOT NULL,
                          status         TEXT NOT NULL,
                          created_at     TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                          updated_at     TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;
-- +goose StatementEnd

