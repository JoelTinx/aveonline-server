CREATE DATABASE aveonlinedb;

CREATE TABLE medicine (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50),
  price DECIMAL(11, 2),
  location VARCHAR(50)
);

CREATE TABLE promotion (
  id SERIAL PRIMARY KEY,
  description VARCHAR(100),
  percentage DECIMAL(5, 2),
  start_date DATE,
  end_date DATE
);

CREATE TABLE invoice (
  id SERIAL PRIMARY KEY,
  creation_date TIMESTAMP,
  total DECIMAL(11, 2)
);
