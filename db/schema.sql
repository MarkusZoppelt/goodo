DROP TABLE IF EXISTS todos;
DROP TABLE IF EXISTS tasks;
CREATE TABLE todos(
    id TEXT PRIMARY KEY,
    name TEXT,
    description TEXT,
    tasks BLOB
);
CREATE TABLE tasks(
    id TEXT PRIMARY KEY,
    name TEXT
);