CREATE DATABASE IF NOT EXISTS ibs_login CHARACTER SET 'UTF8';
USE ibs_login;
CREATE USER IF NOT EXISTS 'ibs'@'localhost' IDENTIFIED BY 'admin';
GRANT ALL ON ibs_login to 'ibs'@'localhost';
