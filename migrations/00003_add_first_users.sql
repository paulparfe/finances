-- +goose Up
-- +goose StatementBegin
INSERT INTO users (id, name, balance) VALUES (1, 'Adam', 111.11);
INSERT INTO users (id, name, balance) VALUES (2, 'Diana', 12.34);

INSERT INTO transactions (id, user_id, recipient_id, amount, transaction_type, created_at)
    VALUES (1, 1, NULL, 123.45, 'deposit', '2025-02-03 11:11:11');
INSERT INTO transactions (id, user_id, recipient_id, amount, transaction_type, created_at)
    VALUES (2, 1, 2, 12.34, 'transfer', '2025-02-04 12:12:12');

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DELETE FROM users;
ALTER SEQUENCE users_id_seq RESTART WITH 1;

DELETE FROM transactions;
ALTER SEQUENCE transactions_id_seq RESTART WITH 1;
-- +goose StatementEnd
