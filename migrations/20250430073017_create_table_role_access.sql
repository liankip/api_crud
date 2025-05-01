-- +goose Up
-- +goose StatementBegin
CREATE TABLE role_access (
    role_id INT,
    access_id INT,
    PRIMARY KEY (role_id, access_id),
    FOREIGN KEY (role_id) REFERENCES role(role_id),
    FOREIGN KEY (access_id) REFERENCES access(access_id)
);
-- +goose StatementEnd
INSERT INTO role_access (role_id, access_id) VALUES
    (1, 1), -- admin: create
    (1, 2), -- admin: read
    (1, 3), -- admin: update
    (1, 4), -- admin: delete
    (2, 1), -- member: create
    (2, 2), -- member: read
    (2, 3), -- member: update
    (3, 2); -- guest: read

-- +goose Down
-- +goose StatementBegin
DROP TABLE role_access
-- +goose StatementEnd
