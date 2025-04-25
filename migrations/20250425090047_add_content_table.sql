
-- +goose Up
CREATE TABLE content (
    id INT NOT NULL AUTO_INCREMENT,
    hash VARCHAR(255),
    sourcecode TEXT,
    language VARCHAR(255),
    PRIMARY KEY (id)
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
