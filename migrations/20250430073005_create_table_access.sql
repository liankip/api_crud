-- +goose Up
-- +goose StatementBegin
CREATE TABLE access (
    access_id INT PRIMARY KEY,
    access_name VARCHAR(10)
);
-- +goose StatementEnd

INSERT INTO access (access_id, access_name) VALUES
    (1, 'create'),
    (2, 'read'),
    (3, 'update'),
    (4, 'delete');

-- +goose Down
-- +goose StatementBegin
DROP TABLE access
-- +goose StatementEnd
