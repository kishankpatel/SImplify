CREATE DATABASE simplify with owner=francium;
GRANT ALL PRIVILEGES ON DATABASE simplify TO francium;

CREATE TABLE books(
  id SERIAL PRIMARY KEY,
  name VARCHAR(50),
  author VARCHAR(120),
  price VARCHAR(120),
  created_by VARCHAR(120),
  created_at DATE NOT NULL DEFAULT CURRENT_DATE
);

-- use database simplify as francium user
\c simplify francium
