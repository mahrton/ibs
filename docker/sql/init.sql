CREATE DATABASE IF NOT EXISTS ibs_login CHARACTER SET 'UTF8';
USE ibs_login;

CREATE TABLE users (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name CHAR(200) NOT NULL,
    password CHAR(100) NOT NULL,
    created_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX USING BTREE (id),
    INDEX USING BTREE (name)
);

CREATE TABLE sessions (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    hash CHAR(200) NOT NULL,
    created_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    INDEX USING BTREE (id),
    INDEX USING BTREE (user_id)
);

CREATE USER IF NOT EXISTS 'ibs'@'localhost' IDENTIFIED BY 'admin';
GRANT ALL ON ibs_login.* to 'ibs'@'localhost';
