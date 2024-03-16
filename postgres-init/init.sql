CREATE DATABASE service_users;
-- switch / use service_users for create users table
\c service_users;
-- Create table for user service
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(225) NOT NULL
);

-- Create database for service-employee
CREATE DATABASE service_employees;
-- switch / use service_employees for create employees table
\c service_employees;

-- Create table for employee service
CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);