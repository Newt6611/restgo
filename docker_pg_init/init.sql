CREATE ROLE m_admin LOGIN PASSWORD 'admin_password' NOSUPERUSER INHERIT NOCREATEDB NOCREATEROLE NOREPLICATION;
CREATE ROLE m_user LOGIN PASSWORD 'user_password' NOSUPERUSER INHERIT NOCREATEDB NOCREATEROLE NOREPLICATION;

CREATE DATABASE bookstore with ENCODING = 'UTF8' LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8' CONNECTION LIMIT = -1 template=template0;

ALTER DATABASE bookstore OWNER TO m_admin;
ALTER DATABASE bookstore SET timezone TO 'UTC';


\connect bookstore;
CREATE TABLE IF NOT EXISTS books (
	id serial PRIMARY KEY,
	title VARCHAR(50) NOT NULL, 
	author VARCHAR(50) NOT NULL,
	created_on TIMESTAMP NOT NULL
);

INSERT INTO books(title, author, created_on)
VALUES('Harry Potter', 'JK Rowling', current_timestamp),
('The Lord of the Rings','J. R. R. Tolkien', current_timestamp);

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO m_admin;

GRANT SELECT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO m_user;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO m_user;