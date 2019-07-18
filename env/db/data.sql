DROP TABLE IF EXISTS todos;

-- Create todo table
CREATE TABLE IF NOT EXISTS todos (
	id int,
	todo_name varchar(30),
	completed boolean
);

INSERT INTO todos VALUES ( 1, 'build MVP API', TRUE);
INSERT INTO todos VALUES ( 2, 'seed database', FALSE);
INSERT INTO todos VALUES ( 3, 'finalize docker compose', FALSE);
INSERT INTO todos VALUES ( 4, 'deploy', FALSE);