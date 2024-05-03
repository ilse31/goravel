CREATE TABLE devices (
    id SERIAL PRIMARY KEY NOT NULL, created_at timestamp NOT NULL, updated_at timestamp NOT NULL, deleted_at timestamp, token VARCHAR(255) NOT NULL, device_id VARCHAR(255) NOT NULL, device_ip VARCHAR(255) NOT NULL, user_id INT NOT NULL UNIQUE, FOREIGN KEY (user_id) REFERENCES users (id) on delete cascade
);