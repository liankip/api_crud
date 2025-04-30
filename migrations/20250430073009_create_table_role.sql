-- +goose Up
-- +goose StatementBegin
CREATE TABLE role (
    role_id INT PRIMARY KEY,
    role_name VARCHAR(10)
);
-- +goose StatementEnd

INSERT INTO role (role_id, role_name) VALUES
    (1, 'admin'),
    (2, 'member'),
    (3, 'guest');

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
