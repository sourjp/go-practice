CREATE TABLE IF NOT EXISTS todo (
    id SERIAL NOT NULL PRIMARY KEY,
    title VARCHAR(20),
    message VARCHAR(255),
    created_at TIMESTAMP,
    finished_at TIMESTAMP
)


